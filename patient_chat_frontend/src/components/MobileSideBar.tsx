import { Home, Hospital, LineChart, MessageCircleHeart, PanelLeft } from "lucide-react"
import { Sheet, SheetContent, SheetTrigger } from "./ui/sheet"
import { Link } from "react-router-dom"
import { Button } from "./ui/button"

export const MobileSideBar = () => {
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
                <LineChart className="h-5 w-5" />
                Settings
              </Link>
            </nav>
          </SheetContent>
        </Sheet>
      </header>
  );
}
