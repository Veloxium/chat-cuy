import { useUserStore } from "@/store/userStore";

export default function Profile() {
   const user = useUserStore((state) => state.user);
   console.log(user.avatar);
   return (
      <div>
         <div className="flex justify-center items-center flex-col h-screen">
            <div className="h-32 w-32 relative">
               <img
                  src={user.avatar}
                  alt="err"
                  crossOrigin="anonymous"
                  referrerPolicy="no-referrer"
               />
            </div>
            <p>{user.created_at}</p>
            <p>{user.username}</p>
            <p>{user.email}</p>
            <p>{user.bio}</p>
         </div>
      </div>
   );
}
