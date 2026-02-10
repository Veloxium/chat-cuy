import { create } from 'zustand'
import { persist, createJSONStorage } from 'zustand/middleware'

type User = {
    id: number
    email: string
    username: string
    bio: string
    avatar: string
    token: string
}

type UserStore = {
   user?: User
   addUser: (user: User) => void
   updateAvatar: (avatar: string) => void
   logout: () => void
}


export const useUserStore = create<UserStore>()(
    persist(
        (set) => (
            {
                user: undefined,
                addUser: (user: User) => set({ user: user }),
                updateAvatar: (avatar: string) => set((state) => ({
                    user: state.user ? { ...state.user, avatar } : undefined,
                })),
                logout: () => set({ user: undefined }),
            }
        ),
        {
            name: 'user-storage',
            storage: createJSONStorage(() => sessionStorage),
        },
    ),
)
