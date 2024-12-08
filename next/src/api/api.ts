import {config} from "@/config/config";

export interface Post {
    id: number;
    slug: string;
    title: string;
    bodyHtml: string;
    postedAt: string;
}

export async function fetchPosts(): Promise<Post[]> {
    const res = await fetch(`${config.BACKEND_URL}/api/posts`).then(res => res.json());
    return res.posts;
}