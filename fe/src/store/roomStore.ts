import { create } from 'zustand'

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


export const useRoomStore = create<RoomStore>()((set) => (
    {
        room: undefined,
        addRoom: (room: RoomType) => set({ room: room }),
        clearRoom: () => set({ room: undefined }),
    }
),

)
