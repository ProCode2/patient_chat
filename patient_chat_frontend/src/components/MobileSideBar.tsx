import { Hospital, LogIn, LogOut, MessageCircleHeart, PanelLeft, Plus, Settings } from "lucide-react"
import { Sheet, SheetContent, SheetTrigger } from "./ui/sheet"
import { Link } from "react-router-dom"
import { Button } from "./ui/button"
import { useMutation } from "@tanstack/react-query";
import { useEffect } from "react"
import { toast } from "./ui/use-toast"

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

export const MobileSideBar = () => {
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
    <header className="sticky top-0 z-30 flex h-14 items-center gap-4 border-b bg-background px-4 sm:static sm:h-auto sm:border-0 sm:bg-transparent sm:px-6">
      <Sheet>
        <SheetTrigger asChild>
          <Button size="icon" variant="outline" className="sm:hidden">
            <PanelLeft className="h-5 w-5" />
            <span className="sr-only">Toggle Menu</span>
          </Button>
        </SheetTrigger>
        <SheetContent side="left" className="sm:max-w-xs">
          <nav className="grid gap-6 text-lg font-medium">
            <Link
              to="/"
              className="group flex h-10 w-10 shrink-0 items-center justify-center gap-2 rounded-full bg-primary text-lg font-semibold text-primary-foreground md:text-base"
            >
              <Hospital className="h-5 w-5 transition-all group-hover:scale-110" />
              <span className="sr-only">Virginia City Hospital</span>
            </Link>

            {
              loggedIn
                ? (
                  <>

                    <Link
                      to="/chat"
                      className="flex items-center gap-4 px-2.5 text-muted-foreground hover:text-foreground"
                    >
                      <MessageCircleHeart className="h-5 w-5" />
                      Chat
                    </Link>
                    <Link
                      to="/settings"
                      className="flex items-center gap-4 px-2.5 text-muted-foreground hover:text-foreground"
                    >
                      <Settings className="h-5 w-5" />
                      Settings
                    </Link>
                    <div
                      className="flex items-center gap-4 px-2.5 text-muted-foreground hover:text-foreground"
                      onClick={() => logoutMutation.mutate()}
                    >
                      <LogOut className="h-5 w-5" />
                      Log out
                    </div>
                  </>
                )
                : (
                  <>

                    <Link
                      to="/login"
                      className="flex items-center gap-4 px-2.5 text-muted-foreground hover:text-foreground"
                    >
                      <LogIn className="h-5 w-5" />
                      Log in
                    </Link>
                    <Link
                      to="/signup"
                      className="flex items-center gap-4 px-2.5 text-muted-foreground hover:text-foreground"
                    >
                      <Plus className="h-5 w-5" />
                      Sign up
                    </Link>
                  </>
                )
            }
          </nav>
        </SheetContent>
      </Sheet>
    </header>
  );
}
