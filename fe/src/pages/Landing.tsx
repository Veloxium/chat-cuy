import { Button } from "@/components/ui/button";
import IndexLayout from "@/layout/IndexLayout";
import { toast } from "sonner";

function LandingPage() {
  return (
    <IndexLayout>
      <div className="flex flex-col items-center justify-center h-screen gap-8">
        <div className="text-center">
          <p className="text-2xl font-forta">Welcome to</p>
          <p className="text-5xl font-forta">Chat Cuy😲</p>
        </div>
        <Button
          variant="default"
          className=""
          onClick={() =>
            toast("Event has been created", {
              description: "Sunday, December 03, 2023 at 9:00 AM",
              action: {
                label: "Undo",
                onClick: () => console.log("Undo"),
              },
            })
          }
        >
          Show Toast
        </Button>
      </div>
    </IndexLayout>
  );
}

export default LandingPage;
