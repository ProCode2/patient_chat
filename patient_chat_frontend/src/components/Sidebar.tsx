import { Link } from "react-router-dom"
import {
  Tooltip,
  TooltipContent,
  TooltipTrigger,
} from "@/components/ui/tooltip"
import { Hospital, Loader, LogIn, LogOut, MessageCircleHeart, Plus, Settings } from "lucide-react"
import { useMutation } from "@tanstack/react-query";
import { useEffect } from "react";
import { toast } from "./ui/use-toast";

const logOut = async () => {
  const res = await fetch("/api/logout", {
    method: "delete",
    headers: {
      "Content-Type": "application/json",
      "Authentication": window.localStorage.getItem("session") || ""
    }
  });
  if (res.status !== 200) {
    throw Error("Can not log out at the moment.")
  }
  window.localStorage.removeItem("session");
  window.location.href = "/";
}

export const SideBar = () => {
  const session = window.localStorage.getItem("session");
  const loggedIn = session != null;
  const logoutMutation = useMutation({ mutationFn: logOut })

  useEffect(() => {
    if (logoutMutation.isError) {
      toast({
        title: "Error while logging out",
        description: "Can not log you out at the moment, Please try again later."
      });
    }
  }, [logoutMutation.isError])

  return (
    <aside className="fixed inset-y-0 left-0 z-10 hidden w-14 flex-col border-r bg-background sm:flex">
      <nav className="h-full flex flex-col items-center gap-4 px-2 sm:py-5">
        <Tooltip>
          <TooltipTrigger asChild>
            <Link
              to="/"
              className="group flex h-9 w-9 shrink-0 items-center justify-center gap-2 rounded-full bg-primary text-lg font-semibold text-primary-foreground md:h-8 md:w-8 md:text-base"
            >
              <Hospital className="h-4 w-4 transition-all group-hover:scale-110" />
              <span className="sr-only">Virginia City Hospital</span>
            </Link>
          </TooltipTrigger>
          <TooltipContent side="right">Virginia City Hospital</TooltipContent>
        </Tooltip>
        {
          loggedIn
            ? (
              <>
                <Tooltip>
                  <TooltipTrigger asChild>
                    <Link
                      to="/chat"
                      className="flex h-9 w-9 items-center justify-center rounded-lg text-muted-foreground transition-colors hover:text-foreground md:h-8 md:w-8"
                    >
                      <MessageCircleHeart className="h-5 w-5" />
                      <span className="sr-only">Chat</span>
                    </Link>
                  </TooltipTrigger>
                  <TooltipContent side="right">Chat</TooltipContent>
                </Tooltip>

                <nav className="mt-auto flex flex-col items-center gap-4 px-2 sm:py-5">
                  <Tooltip>
                    <TooltipTrigger asChild>
                      <>
                        {logoutMutation.isPending ? <Loader className="h-5 w-5 animate-spin" /> : <LogOut onClick={() => logoutMutation.mutate()} className="h-5 w-5" />}                        <span className="sr-only">Log out</span>
                      </>
                    </TooltipTrigger>
                    <TooltipContent side="right">Log out</TooltipContent>
                  </Tooltip>
                  <Tooltip>
                    <TooltipTrigger asChild>
                      <Link
                        to="/settings"
                        className="flex h-9 w-9 items-center justify-center rounded-lg text-muted-foreground transition-colors hover:text-foreground md:h-8 md:w-8"
                      >
                        <Settings className="h-5 w-5" />
                        <span className="sr-only">Settings</span>
                      </Link>
                    </TooltipTrigger>
                    <TooltipContent side="right">Settings</TooltipContent>
                  </Tooltip>
                </nav>
              </>
            )
            :
            <>
              <Tooltip>
                <TooltipTrigger asChild>
                  <Link
                    to="/login"
                    className="flex h-9 w-9 items-center justify-center rounded-lg text-muted-foreground transition-colors hover:text-foreground md:h-8 md:w-8"
                  >
                    <LogIn className="h-5 w-5" />
                    <span className="sr-only">Log in</span>
                  </Link>
                </TooltipTrigger>
                <TooltipContent side="right">Log in</TooltipContent>
              </Tooltip>

              <Tooltip>
                <TooltipTrigger asChild>
                  <Link
                    to="/signup"
                    className="flex h-9 w-9 items-center justify-center rounded-lg text-muted-foreground transition-colors hover:text-foreground md:h-8 md:w-8"
                  >
                    <Plus className="h-5 w-5" />
                    <span className="sr-only">Sign up</span>
                  </Link>
                </TooltipTrigger>
                <TooltipContent side="right">Sign up</TooltipContent>
              </Tooltip>
            </>
        }
      </nav>
    </aside>
  )
}
