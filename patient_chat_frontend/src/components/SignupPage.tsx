import { Button } from "@/components/ui/button"
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card"
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select"
import { Input } from "@/components/ui/input"
import { Label } from "@/components/ui/label"
import { useEffect, useState } from "react"
import { Link } from "react-router-dom"
import {
  useQuery,
  useMutation
} from '@tanstack/react-query';
import { useToast } from "./ui/use-toast"
import { Loader } from "lucide-react"

interface ISignupState {
  phoneNumber: string;
  otp: string;
  firstName: string;
  lastName: string;
  doctorId: string;
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

const handleSignupSubmit = async (data: ISignupState) => {
  if (data.otp === "" || data.phoneNumber.length != 10 || data.firstName.length < 3 || data.lastName.length < 3) {
    throw Error("Please fill all the fields before submitting.");
    
  }
  const res = await fetch("/api/signup", {
    method: "post",
    headers: {
      "Content-Type": "application/json"
    },
    body: JSON.stringify({ ...data, patientName: `${data.firstName} ${data.lastName}` })
  });
  if (res.status != 200) {
    throw Error("Can not sign in at the moment, please try again later");
  }

  const session = await res.json();

  console.log(session);
  localStorage.setItem("session", session.sessionId);
  window.location.href = "/chat";
}

export const SignupPage = () => {
  const [state, setState] = useState<ISignupState>({
    phoneNumber: "",
    firstName: "",
    lastName: "",
    otp: "",
    doctorId: ""
  });

  const { toast } = useToast();

  // Queries
  const { isFetching, isSuccess, isError, data, error } = useQuery<DDoc[]>({
    queryKey: ['allDoctors'], queryFn: getAllDoctors
  })
  const signupMutation = useMutation({
    mutationFn: handleSignupSubmit
  })

  useEffect(() => {
    console.log(signupMutation.isSuccess)
    if (isError) {
      toast({
        title: "Error while getting doctors",
        description: error.message
      });
    }

    if (signupMutation.isSuccess) {
      toast({
        title: "Welcome!",
        description: "You have been signed up!",
      });
    }

    if (signupMutation.isError) {
      toast({
        title: "Error while signing up",
        description: signupMutation.error?.message
      });
    }
  }, [isError, error, signupMutation.isSuccess, signupMutation.isError, signupMutation.error])



  return (
    <Card className="mx-auto max-w-sm">
      <CardHeader>
        <CardTitle className="text-xl">Sign Up</CardTitle>
        <CardDescription>
          Enter your information to create an account
        </CardDescription>
      </CardHeader>
      <CardContent>
        <div className="grid gap-4">
          <div className="grid grid-cols-2 gap-4">
            <div className="grid gap-2">
              <Label htmlFor="first-name">First name</Label>
              <Input id="first-name" placeholder="Max" required onChange={(e) => setState((prev) => ({ ...prev, firstName: e.target.value }))} />
            </div>
            <div className="grid gap-2">
              <Label htmlFor="last-name">Last name</Label>
              <Input id="last-name" placeholder="Robinson" required onChange={(e) => setState((prev) => ({ ...prev, lastName: e.target.value }))} />
            </div>
          </div>
          <div className="grid gap-2">
            <Label htmlFor="phone">Phone</Label>
            <Input
              id="phone"
              type="phone"
              placeholder="729*******"
              value={state.phoneNumber}
              onChange={(e) => setState((prev) => ({ ...prev, phoneNumber: e.target.value }))}
              required
            />
          </div>

          <div className="grid gap-2">
            <Label htmlFor="doctor">Doctor</Label>
            {
              isFetching
                ? <span>Loading all doctors</span>
                :
                <Select onValueChange={(doctor) => setState((prev) => ({ ...prev, doctorId: doctor }))}>
                  <SelectTrigger className="w-[180px]">
                    <SelectValue placeholder="Doctor" />
                  </SelectTrigger>
                  <SelectContent>
                    {
                      isSuccess
                        ? data?.map(doc => <SelectItem key={doc.id} value={doc.id}>{doc.name}</SelectItem>)
                        : null
                    }
                  </SelectContent>
                </Select>
            }
          </div>
          <div className="grid gap-2">
            <Label htmlFor="password">Otp</Label>
            <Input id="otp" type="text" onChange={(e) => setState((prev) => ({ ...prev, otp: e.target.value }))} />
          </div>
          <Button type="submit" className="w-full" onClick={() => signupMutation.mutate(state)}>
            <span className="mr-2">Create an account</span>
            {signupMutation.isPending ? <span><Loader className="animate-spin" /></span> : null}
          </Button>
        </div>
        <div className="mt-4 text-center text-sm">
          Already have an account?{" "}
          <Link to="/login" className="underline">
            Log in
          </Link>
        </div>
      </CardContent>
    </Card>
  )
}
