import { Outlet } from "react-router-dom"
import { SideBar } from "./Sidebar"
import { MobileSideBar } from "./MobileSideBar"

export const AppLayout = () => {
  return (
    <div className="flex min-h-screen w-full flex-col bg-muted/40">
      <SideBar />
      <div className="flex flex-col sm:gap-4 sm:py-4 sm:pl-14">
        <MobileSideBar />
        <Outlet />
      </div>
    </div>
  )
}
