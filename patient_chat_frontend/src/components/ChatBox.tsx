import { useMutation, useQuery } from "@tanstack/react-query";
import { useParams } from "react-router-dom";
import { Label } from "./ui/label";
import { Button } from "./ui/button";
import { CornerDownLeft, Loader } from "lucide-react";
import { Textarea } from "./ui/textarea";
import { useEffect, useState } from "react";
import { toast } from "./ui/use-toast";

interface IChatType {
  id: string;
  patientId: string;
  doctorId: string;
  threadId: string;
  query: string;
  response: string;
  time: string;
}


const getChat = async (threadId: string) => {
  const res = await fetch("/api/patient/chats/" + threadId, {
    headers: {
      "Content-Type": "application/json",
      "Authentication": window.localStorage.getItem("session") || ""
    }
  });

  if (res.status !== 200) {
    throw Error("Something went wrong while getting chats");
  }

  const chats = await res.json();
  return chats;
}

const sendChat = async ({ threadId, query }: { threadId: string, query: string }) => {
  const res = await fetch("/api/patient/chats", {
    method: "post",
    headers: {
      "Content-Type": "application/json",
      "Authentication": window.localStorage.getItem("session") || ""
    },
    body: JSON.stringify({ threadId, query })
  });

  if (res.status != 200) {
    throw Error("Something went wrong while sending that message, Pleasse try again!");
  }

  const chat = await res.json();

  window.location.href = "/chat/" + chat.threadId;
}

const UserAndBotMessage = ({ ch }: { ch: IChatType }) => {

  return (
    <>
      <div key={ch.id} className="w-full flex flex-col gap-2 justify-center items-start text-left my-2 border border-slate-300 rounded-md p-2">
        <div className="flex justify-center items-center ">
          <span className="bg-primary rounded-full w-10 h-10 text-sm text-white flex justify-center items-center">
            You
          </span>
          <span className="ml-2">
            {ch.query}
          </span>
        </div>
        <p className="text-slate-400 font-mono text-sm ml-auto">{ch.time}</p>
      </div>
      <div key={ch.id} className="w-full flex flex-col gap-2 justify-center items-start text-left my-2 border border-slate-300 rounded-md p-2">
        <div className="flex justify-center items-center ">
          <span className="bg-primary rounded-full w-10 h-10 text-sm text-white flex justify-center items-center">
            Bot
          </span>
          <span className="ml-2">
            {ch.response}
          </span>
        </div>
        <p className="text-slate-400 font-mono text-sm ml-auto">{ch.time}</p>
      </div>
    </>
  )
}


export const ChatBox = () => {
  const { threadId } = useParams();
  const [query, setQuery] = useState("");

  const getThread = useQuery<IChatType[]>({ queryKey: ["get-chats", threadId], queryFn: ({ queryKey }) => getChat(queryKey[1] as string), enabled: !!threadId })

  const addChat = useMutation({ mutationFn: sendChat })

  useEffect(() => {
    if (addChat.isError) {
      toast({
        title: "Error while sending chat",
        description: addChat.error.message
      });
    }
  }, [addChat.isError, addChat.error]);

  return (
    <div className="pt-4 border-1 border-primary-400 rounded overflow-auto">
      {
        !getThread.data?.length ? <p className="font-bold text-xl text-slate-500 my-4">Start asking your query</p> : null
      }
      {
        getThread?.data?.map(ch => {
          return <UserAndBotMessage ch={ch} />
        })
      }
      <div>
        <div className="relative overflow-hidden rounded-lg border bg-background focus-within:ring-1 focus-within:ring-ring">
          <Label htmlFor="message" className="sr-only">
            Message
          </Label>
          <Textarea
            value={query}
            onChange={(e) => setQuery(e.target.value)}
            id="message"
            placeholder="Type your message here..."
            className="min-h-12 resize-none border-0 p-3 shadow-none focus-visible:ring-0"
          />
          <div className="flex items-center p-3 pt-0">
            <Button type="submit" size="sm" className="ml-auto gap-1.5" disabled={addChat.isPending} onClick={() => addChat.mutate({threadId: threadId || "", query})}>
              Send Message
              <CornerDownLeft className="size-3.5" />
              {addChat.isPending ? <Loader className="w-5 h-5 animate-spin" /> : null}
            </Button>
          </div>
        </div>
      </div>
    </div>
  )
}
