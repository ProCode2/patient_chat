import { Hospital, MessageCircleDashed } from "lucide-react";
import { Link } from "react-router-dom";

export const HomePage = () => {
  return (
    <main className="w-full h-full">
      <div className="w-full h-full">
        <div className="w-full h-full flex flex-col justify-center items-center">
          <div className="bg-primary rounded-full hover:shadow-lg py-2 px-4 hover:scale-110 delat-350 ease-in-out transition-all mb-2">
            <Hospital className="w-12 h-13 text-white " />
          </div>
          <h3 className="font-semibold text-base sm:text-xl md:text-2xl">Welcome to Virginia City Hospital</h3>
          <p className="font-bold text-xl sm:text-2xl md:text-4xl uppercase tracking-wider mt-4 mb-2 text-center">Answering your medical queries</p>
          <span className="text-center max-w-2xl sm:max-w-3xl md:max-w-4xl">Ofcourse all of this is made up. There is no AI inside it is just a take home assesment the AI module is simulated and do not use it for real needs. Consult a doctor if you need help!</span>
          <Link to="/chat">
            <div className="flex items-center gap-2 bg-primary rounded-full hover:shadow-lg py-2 px-4 hover:scale-110 delat-350 ease-in-out transition-all mt-2">
              <MessageCircleDashed className="text-white" />
              <span className="text-white text-sm">Chat with AI</span>
            </div>
          </Link>
        </div>
      </div>
    </main>
  );
}
