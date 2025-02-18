import { create } from "zustand";
import { persist, createJSONStorage } from "zustand/middleware";

export type User = {
   id: string;
   email: string;
   username: string;
   accessToken: string;
   avatar: string;
   bio: string;
   created_at: string;
};

export interface ResponseInfo {
   success: boolean;
   statusCode: number;
   message: string;
   data: User;
}

type UserStore = {
   authenticated: boolean;
   setAuthenticated: (auth: boolean) => void;
   user: User;
   setUser: (user: User) => void;
   // user?: User;
   // addUser: (user: User) => void;
   // logout: () => void;
};

export const useUserStore = create<UserStore>()(
   persist(
      (set) => ({
         user: {
            id: "",
            email: "",
            username: "",
            accessToken: "",
            avatar: "",
            bio: "",
            created_at: "",
         },
         authenticated: false,
         setAuthenticated: (authenticated: boolean) =>
            set({ authenticated: authenticated }),
         setUser: (user: User) => set({ user: user }),
         // addUser: (user: User) => set({ user: user }),
         // logout: () => set({ user: undefined }),
      }),
      {
         name: "user-storage",
         storage: createJSONStorage(() => sessionStorage), // (optional) by default, 'localStorage' is used
      },
   ),
);
