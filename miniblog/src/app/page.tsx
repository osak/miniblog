interface Post {
    id: number;
    slug: string;
    title: string;
    bodyHtml: string;
    postedAt: string;
}

export default async function Index() {
    const res = await fetch('http://localhost:8081/api/posts').then(res => res.json());
    const posts: Post[] = res.posts;
    return (
        <div>
            <ul>
                {posts.map(post => (
                    <div className="post" key={post.id}>
                        <div className="post-header">
                            <h2>{post.title}</h2>
                            <div className="post-date">{post.postedAt}</div>
                        </div>
                        <div dangerouslySetInnerHTML={{__html: post.bodyHtml}} />
                    </div>
                ))}
            </ul>
        </div>
    )
}
