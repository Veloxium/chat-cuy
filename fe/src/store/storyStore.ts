import { StoryItem } from '@/data/storyitem'
import { create } from 'zustand'

type StoryUrlStore = {
    story?: StoryItem
    addStory: (story: StoryItem) => void
    clearStory: () => void
}

export const useStoryStore = create<StoryUrlStore>((set) => ({
    story: undefined,
    addStory: (story: StoryItem) => setTimeout(() => set({ story: story }), 0),
    clearStory: () => setTimeout(() => set({ story: undefined }), 0),
}))
