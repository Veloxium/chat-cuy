import { number } from "zod";

const dataStoryItem = [
    {
        id: 1,
        username: "Sigma",
        avatar: "https://api.dicebear.com/9.x/adventurer/svg?seed=Jude",
        text: "Kelass",
        url: "https://i.pinimg.com/736x/5b/8e/b5/5b8eb524dc01b7201d50df40b4034f7e.jpg",
        created_at: "2021-01-01T11:56:00Z",
    },
    {
        id: 2,
        username: "Alpha",
        avatar: "https://api.dicebear.com/9.x/adventurer/svg?seed=Brian",
        text: "Hello World",
        url: "https://i.pinimg.com/736x/26/de/8c/26de8cda547928f09fe1e201434b784d.jpg",
        created_at: "2021-01-02T12:00:00Z",
    },
    {
        id: 3,
        username: "Beta",
        avatar: "https://api.dicebear.com/9.x/adventurer/svg?seed=Aiden",
        text: "Good Morning",
        url: "https://i.pinimg.com/736x/01/14/46/01144668c777a5819513e20cbc1e6b6a.jpg",
        created_at: "2021-01-03T08:30:00Z",
    },
    {
        id: 4,
        username: "Gamma",
        avatar: "https://api.dicebear.com/9.x/adventurer/svg?seed=Leo",
        text: "How are you?",
        url: "https://i.pinimg.com/736x/65/7e/21/657e21cea3458a5d02001f9ead192a85.jpg",
        created_at: "2021-01-04T09:45:00Z",
    },
    {
        id: 5,
        username: "Delta",
        avatar: "https://api.dicebear.com/9.x/adventurer/svg?seed=Mason",
        text: "Nice to meet you",
        url: "https://i.pinimg.com/736x/d1/9b/bf/d19bbf9120eea4ddf2e6c8f40906ac74.jpg",
        created_at: "2021-01-05T10:15:00Z",
    },
    {
        id: 6,
        username: "Epsilon",
        avatar: "https://api.dicebear.com/9.x/adventurer/svg?seed=Maria",
        text: "Good Night",
        url: "https://i.pinimg.com/736x/09/22/68/0922688f238b9d800ccc911326772d8b.jpg",
        created_at: "2021-01-06T22:00:00Z",
    },
    {
        id: 7,
        username: "Zeta",
        avatar: "https://api.dicebear.com/9.x/adventurer/svg?seed=Jocelyn",
        text: "Babay",
        url: "https://i.pinimg.com/736x/fb/3a/b3/fb3ab3617b9718a2deec86a35d0cf9b2.jpg",
        created_at: "2021-01-07T23:59:00Z",
    },
]

type StoryItem = {
    id: number;
    username: string;
    avatar: string;
    text: string;
    url: string;
    created_at: string;
}

export { dataStoryItem };
export type { StoryItem };
