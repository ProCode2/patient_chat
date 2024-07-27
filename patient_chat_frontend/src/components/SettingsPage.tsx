import { useMutation, useQuery } from "@tanstack/react-query";
import { Button } from "./ui/button";
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "./ui/card";
import { Input } from "./ui/input";
import { useEffect, useState } from "react";
import { toast } from "./ui/use-toast";
import { Select, SelectContent, SelectItem, SelectTrigger, SelectValue } from "./ui/select";

const getPatient = async () => {
  const res = await fetch("/api/patient", {
    headers: {
      "Content-Type": "application/json",
      "Authentication": window.localStorage.getItem("session") || ""
    }
  });

  if (res.status !== 200) {
    throw Error("Something went wrong while fetching patient data, Please try again!");
  }

  const patient = await res.json()
  return patient;
}

const updatePatient = async (data: IInputState) => {
  const res = await fetch("/api/patient", {
    method: "put",
    headers: {
      "Content-Type": "application/json",
      "Authentication": window.localStorage.getItem("session") || ""
    },
    body: JSON.stringify(data)
  });

  if (res.status !== 200) {
    throw Error("Something went wrong while updatig patient data, Please try again!");
  }

  // window.location.reload()
}

interface IUser {
  id: string;
  name: string;
  role: string;
  phone: string;
}

interface IPatient {
  id: string;
  userId: string;
  docId: string;
  medicalHistory: string;
}

interface IPatientUser {
  user: IUser,
  patient: IPatient
}

interface IInputState {
  name: string;
  docId: string;
  medicalHistory: string;
}


interface DDoc {
  name: string;
  id: string;
}

const getAllDoctors = async () => {
  const res = await fetch("/api/docs");
  const docs = await res.json();
  return docs;
}

export const SettingsPage = () => {
  const patientQuery = useQuery<IPatientUser>({ queryKey: ["get-patient-user"], queryFn: getPatient })
  const [state, setState] = useState<IInputState>({
    name: patientQuery.data?.user.name || "",
    docId: patientQuery.data?.patient.docId || "",
    medicalHistory: "",
  });
  const getDoc = useQuery<DDoc[]>({
    queryKey: ['allDoctors'], queryFn: getAllDoctors
  });
  const editPatient = useMutation({mutationFn: updatePatient})

  useEffect(() => {
    if (patientQuery.isError) {
      toast({
        title: "Error while getting patient data",
        description: patientQuery.error.message
      });
    }
    setState({
      name: patientQuery.data?.user.name || "",
      docId: patientQuery.data?.patient.docId || "",
      medicalHistory: "",
    })
    console.log(patientQuery.data);
  }, [patientQuery.isError, patientQuery.error, patientQuery.data])

  return (
    <div className="mx-auto grid w-full max-w-6xl items-start gap-6 md:grid-cols-[180px_1fr] lg:grid-cols-[250px_1fr]">
      <div className="grid gap-6 w-screen max-w-2xl sm:max-w-3xl py-3 px-2">
        <Button onClick={() => editPatient.mutate(state)}>Save</Button>
        <Card>
          <CardHeader>
            <CardTitle>Patient Name</CardTitle>
            <CardDescription>
              Enter you legal name as in your Aadhar card
            </CardDescription>
          </CardHeader>
          <CardContent>
            <Input
              placeholder={patientQuery.isPending ? "Loading Name" : "Enter Your Name"}
              value={state.name}
              onChange={(e) => setState((prev) => ({ ...prev, name: e.target.value }))}
            />
          </CardContent>
        </Card>
        <Card>
          <CardHeader>
            <CardTitle>Medical History</CardTitle>
            <CardDescription>
              A history of your previous medical diagnosis
            </CardDescription>
          </CardHeader>
          <CardContent>
            <div className="py-2">
              <ul>
                {
                  patientQuery.data?.patient.medicalHistory ? patientQuery.data?.patient.medicalHistory.split(",").map(diagnosis => <li className="text-sm border border-primary rounded-md my-3 p-2">{diagnosis}</li>) : null
                }
              </ul>
            </div>
            <Input
              placeholder="Add a new diagnosis status"
              value={state.medicalHistory}
              onChange={(e) => setState((prev) => ({ ...prev, medicalHistory: e.target.value }))}
            />
          </CardContent>
        </Card>
        <Card>
          <CardHeader>
            <CardTitle>Your Doctor</CardTitle>
            <CardDescription>
              Linked Doctor Information
            </CardDescription>
          </CardHeader>
          <CardContent>
            <Select onValueChange={(doctor) => setState((prev) => ({ ...prev, docId: doctor }))} value={state.docId}>
              <SelectTrigger className="w-[180px]">
                <SelectValue placeholder="Doctor" />
              </SelectTrigger>
              <SelectContent>
                {
                  getDoc.isSuccess
                    ? getDoc.data?.map(doc => <SelectItem key={doc.id} value={doc.id}>{doc.name}</SelectItem>)
                    : null
                }
              </SelectContent>
            </Select>
          </CardContent>
        </Card>
      </div>
    </div>
  );
}
