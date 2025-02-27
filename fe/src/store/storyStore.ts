import { create } from 'zustand'

export type StoryUrlType = {
    id: number
    username: string
    url: string
}

type StoryUrlStore = {
    story?: StoryUrlType
    addStory: (story: StoryUrlType) => void
    clearStory: () => void
}

export const useStoryStore = create<StoryUrlStore>((set) => ({
    story: undefined,
    addStory: (story: StoryUrlType) => setTimeout(() => set({ story: story }), 0),
    clearStory: () => setTimeout(() => set({ story: undefined }), 0),
}))
