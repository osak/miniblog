import {fetchPosts} from "@/api/api";


export default async function Index() {
    const posts = await fetchPosts();
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
