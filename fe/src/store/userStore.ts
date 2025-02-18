import { create } from 'zustand'
import { persist, createJSONStorage } from 'zustand/middleware'

type User = {
    id: number
    email: string
    username: string
    bio: string
}

type UserStore = {
   user?: User
   addUser: (user: User) => void
   logout: () => void
}


export const useUserStore = create<UserStore>()(
    persist(
        (set) => (
            {
                user: undefined,
                addUser: (user: User) => set({ user: user }),
                logout: () => set({ user: undefined }),
            }
        ),
        {
            name: 'user-storage',
            storage: createJSONStorage(() => sessionStorage), // (optional) by default, 'localStorage' is used
        },
    ),
)
