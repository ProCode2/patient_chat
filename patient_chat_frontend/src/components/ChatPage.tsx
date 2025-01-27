import { Outlet, useParams } from "react-router-dom";
import { ChatList } from "./ChatList";
import { ChatBox } from "./ChatBox";

export const ChatPage = () => {
  const {threadId } = useParams();

  
  return (
    <main className="grid flex-1 items-start gap-4 p-4 sm:px-6 sm:py-0 md:gap-8 lg:grid-cols-3 xl:grid-cols-3">
      <div className="grid auto-rows-max items-start gap-4 md:gap-8 lg:col-span-2">
        <div className="">
          <ChatList />
          {!threadId ? <ChatBox /> : null}
          <Outlet />
        </div>
      </div>
    </main>
  );
}
