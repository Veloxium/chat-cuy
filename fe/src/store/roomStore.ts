import { create } from 'zustand'
import { persist, createJSONStorage } from 'zustand/middleware'

export type RoomType = {
    id: number
    avatar: string
    username: string
    bio: string
}

type RoomStore = {
    room?: RoomType
    addRoom: (room: RoomType) => void
    clearRoom: () => void
}


export const useRoomStore = create<RoomStore>()(
    persist(
        (set) => (
            {
                room: undefined,
                addRoom: (room: RoomType) => set({ room: room }),
                clearRoom: () => set({ room: undefined }),
            }
        ),
        {
            name: 'room-storage',
            storage: createJSONStorage(() => localStorage),
        },
    ),
)
