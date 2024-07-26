import { Button } from "@/components/ui/button"
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card"
import { Input } from "@/components/ui/input"
import { Label } from "@/components/ui/label"
import { useEffect, useState } from "react"
import { Link } from "react-router-dom"
import {
  useMutation
} from '@tanstack/react-query';
import { toast } from "./ui/use-toast"
import { Loader } from "lucide-react"

interface ILoginState {
  phone: string;
  otp: string;
}

const handleLoginSubmit = async (data: ILoginState) => {
  if (data.phone.length != 10 || data.otp.length != 6) {
    throw Error("Please fill all the fields before submitting.");

  }
  const res = await fetch("/api/login", {
    method: "post",
    headers: {
      "Content-Type": "application/json"
    },
    body: JSON.stringify(data)
  });
  if (res.status != 200) {
    throw Error("Can not log in at the moment, please try again later");
  }

  const session = await res.json();

  console.log(session);
  localStorage.setItem("session", session.sessionId);
  window.location.href = "/chat";
}

export const LoginPage = () => {
  const [state, setState] = useState<ILoginState>({
    phone: "",
    otp: ""
  });

  const loginMutation = useMutation({mutationFn: handleLoginSubmit})

  useEffect(() => {
    if (loginMutation.isError) {
      toast({
        title: "Error while loggin in",
        description: loginMutation.error.message
      });
    }

    if (loginMutation.isSuccess) {
      toast({
        title: "Welcome",
        description: "Successfully Logged in!"
      });
    }

  }, [loginMutation.isSuccess, loginMutation.isError, loginMutation.isError]); 

  return (
    <Card className="mx-auto max-w-sm">
      <CardHeader>
        <CardTitle className="text-2xl">Login</CardTitle>
        <CardDescription>
          Enter your email below to login to your account
        </CardDescription>
      </CardHeader>
      <CardContent>
        <div className="grid gap-4">
          <div className="grid gap-2">
            <Label htmlFor="phone">Phone</Label>
            <Input
              id="phone"
              type="phone"
              placeholder="728*******."
              required
              value={state.phone}
              onChange={(e) => setState((prev) => ({...prev, phone: e.target.value}))}
            />
          </div>
          <div className="grid gap-2">
            <div className="flex items-center">
              <Label htmlFor="password">Password</Label>
            </div>
            <Input id="otp" value={state.otp} type="text" placeholder="123456" required onChange={(e) => setState((prev) => ({...prev, otp: e.target.value}))} />
          </div>
          <Button type="submit" className="w-full" onClick={() => loginMutation.mutate(state)}>
            <span>Login</span>
            {loginMutation.isPending ? <span><Loader className="animate-spin" /></span> : null}
          </Button>
        </div>
        <div className="mt-4 text-center text-sm">
          Don&apos;t have an account?{" "}
          <Link to="/signup" className="underline">
            Sign up
          </Link>
        </div>
      </CardContent>
    </Card>
  )
}
