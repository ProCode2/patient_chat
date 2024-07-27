import { Sheet, SheetContent, SheetTrigger } from "./ui/sheet";
import { Button } from "./ui/button";
import { MessageCircleDashed, PanelRight } from "lucide-react";
import { useQuery } from "@tanstack/react-query";
import { Link } from "react-router-dom";

interface IChatType {
  id: string;
  patientId: string;
  doctorId: string;
  threadId: string;
  query: string;
  response: string;
  time: string;
}

const getChats = async () => {
  const res = await fetch("/api/patient/chats", {
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

export const ChatList = () => {
  const getChat = useQuery<IChatType[]>({ queryKey: ["get-chats"], queryFn: getChats })
  return (
    <header className="bg-background sm:static sm:h-auto sm:border-0 sm:bg-transparent">
      <Sheet>
        <SheetTrigger>
          <Button variant="outline" className="">
            <PanelRight className="h-5 w-5 mr-2" />
            <span>View Chats</span>
          </Button>
        </SheetTrigger>
        <SheetContent>
          <div className="pt-6">
            {
              !getChat.data?.length ? <span className="font-semibold text-slate-400 text-xl">No chats found</span> : null
            }

            {
              getChat.data?.map(ch => {
                return <Link
                  key={ch.id}
                  to={`/chat/${ch.threadId}`}
                  className="flex items-center gap-4 px-2.5 text-muted-foreground hover:text-foreground"
                >
                  <MessageCircleDashed className="h-5 w-5" />
                  {`${ch.query.slice(0, 15)}...`}
                </Link>
              })
            }
          </div>
        </SheetContent>
      </Sheet>
    </header>
  );
}
