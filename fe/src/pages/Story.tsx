import ZBackground from "@/components/custom/zbackground";
import SideBarLayout from "@/layout/SideBarLayout";
import { FaPlus } from "react-icons/fa6";

function StoryPage() {
  return (
    <SideBarLayout>
      <p className="text-4xl font-forta">Stories</p>
      <div className="grid grid-cols-2 gap-4 mt-4">
        <div className="relative overflow-hidden z-10 w-full aspect-square place-content-center place-items-center text-center rounded-md border-2 border-dashed border-zprimary">
          <div className="absolute -z-10 w-full h-full top-0">
            <ZBackground />
          </div>
          <FaPlus size={26} className="text-white" />
        </div>
        {Array.from({ length: 9 }).map((_, i) => (
          <div
            key={i}
            className="w-full aspect-square place-content-center place-items-center text-center rounded-md bg-yellow-500"
          >
            Story {i + 1}
          </div>
        ))}
      </div>
    </SideBarLayout>
  );
}

export default StoryPage;
