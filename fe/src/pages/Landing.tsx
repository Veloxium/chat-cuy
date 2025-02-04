import ZBackground from "@/components/custom/zbackground";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Separator } from "@/components/ui/separator";
import IndexLayout from "@/layout/IndexLayout";
import { FaGoogle } from "react-icons/fa";
import { CiCircleQuestion } from "react-icons/ci";
import { toast } from "sonner";
import { Link } from "react-router-dom";

function LandingPage() {
  return (
    <IndexLayout>
      <ZBackground />
      <div className="min-h-screen flex items-center justify-center">
        <div className="bg-white text-zprimary flex flex-col items-center justify-center gap-4 py-8 px-4 w-full max-w-lg rounded-lg">
          <div className="text-center mb-6">
            <p className="text-2xl font-forta">Welcome to</p>
            <p className="text-5xl font-forta">Chat Cuy😲</p>
            <p className="font-semibold mt-2 text-lg">
              Chat dengan siapa pun yang kamu mau
            </p>
          </div>
          <form className="text-black flex flex-col gap-4 items-left w-full max-w-sm text-left">
            <p className="-mb-2">Email</p>
            <Input placeholder="Masukan Email" />
            <p className="-mb-2">Password</p>
            <Input placeholder="Masukan Password" />
            <Link to="/home">
              <button className="flex w-full max-w-sm gap-2 items-center justify-center cursor-pointer">
                <p className="text-lg border-2 border-zprimary bg-zprimary hover:bg-hprimary text-white font-medium rounded-md px-4 py-2 w-full ease-in duration-300">
                  Login Cuy
                </p>
              </button>
            </Link>
          </form>
          <Separator className="w-full max-w-sm" />
          <div className="flex w-full gap-4 items-center justify-center">
            <button
              className="flex w-full max-w-sm gap-2 group items-center justify-center cursor-pointer"
              // onClick={() =>
              //   toast("Event has been created", {
              //     description: "Sunday, December 03, 2023 at 9:00 AM",
              //     action: {
              //       label: "Undo",
              //       onClick: () => console.log("Undo"),
              //     },
              //   })
              // }
            >
              <p className="bg-white border-2 border-zprimary group-hover:bg-hprimary text-zprimary group-hover:text-white font-medium rounded-md px-4 py-2 w-full ease-in duration-300">
                With Google
              </p>
              <div className="bg-white border-2 border-zprimary group-hover:bg-hprimary rounded-md px-2 py-2 ease-in duration-300">
                <FaGoogle
                  size={24}
                  className="text-zprimary group-hover:text-white"
                />
              </div>
            </button>
          </div>
          <div className="mt-4 flex items-center text-sm gap-1 hover:underline justify-end w-full max-w-sm">
            <p>Our Repository</p>
            <CiCircleQuestion size={16} />
          </div>
        </div>
      </div>
    </IndexLayout>
  );
}

export default LandingPage;
