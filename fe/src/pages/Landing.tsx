import { Button } from "@/components/ui/button";
import { toast } from "sonner";

function LandingPage() {
  return (
    <div className="flex flex-col items-center justify-center h-screen gap-8">
      <p className="text-5xl font-forta">LandingPage</p>
      <Button
        variant="default"
        className="bg-green-400 cursor-pointer"
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
  );
}

export default LandingPage;
