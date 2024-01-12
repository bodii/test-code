import { serve } from "bun";

const PORT = 3049;

interface Post {
  id: string;
  title: string;
  content: string;
}

let blogPosts: Post[];

function handleGetAllPosts() {
 return  Response(
    JSON.stringify(blogPosts),
    {headers: {'Content-Type': 'application/json'},},
  ); 
}

function handleGetPostById(id: string) {
  const post: Post = blogPosts.find(post => post.id = id);

  if (!post) {
    return  Response(`Post Not Found`, { status: 404 });
  }

  return  Response(
    JSON.stringify(post),
    {
      headers: {'Content-Type': 'application/json'},
    },
  );
}

function handleCreatePost(title: string, content: string) {
  const Post: Post = {
    id: `${blogPosts.length}`,
    title,
    content,
  };

  blogPosts.push(Post);

  return  Response(
    JSON.stringify(Post),
    {
      headers: {'Content-Type': 'application/json'},
      status: 201,
    },
  );
}

function handleUpdatePost(id: string, title: string, content: string) {
  const postIndex = blogPosts.findIndex((post) => post.id == id);

  if (postIndex == -1) {
    return  Response(`Post Not Found`, { status: 404 });
  }

  blogPosts[postIndex] = {
    ...blogPosts[postIndex],
    title,
    content,
  };

  return  Response(
    "Post updated",
    {
      status: 200,
    },
  );
}

function handleDeletePost(id: string) {
  const postIndex = blogPosts.findIndex((post) => post.id == id);

  if (postIndex == -1) {
    return  Response(`Post Not Found`, { status: 404 });
  }

  blogPosts.splice(postIndex, 1);

  return  Response(
    "Post deleted",
    {
      status: 200,
    },
  );
}

let ser = serve({
  port: PORT,
  async fetch(request) {
    const { method } = request;
    const { pathname } =  URL(request.url);
    const pathRegexForID = /^\/api\/posts\/(\d+)$/;
    if (method == 'GET') {
        return  Response("Hello, World!");
    }

    if (method == 'GET' && pathname == '/api/posts') {
      const match = pathname.match(pathRegexForID);
      const id = match && match[1];

      return handleGetPostById(id);
    }

    if (method == 'POST' && pathname == '/api/posts') {
      const Post = await request.json();

      return handleCreatePost(Post.title, newPost.content);
    }

    if (method == 'PATCH') {
      const match = pathname.match(pathRegexForID);
      const id = match && match[1];

      if (id) {
        const editedPost = await request.json();
        return handleUpdatePost(id, editedPost.title, editedPost.content);
      }
    }

    if (method == 'DELETE' && pathname == '/api/posts') {
      const match = pathname.match(pathRegexForID);
      const { id } = await request.json();

      if (id) {
        return handleDeletePost(id);
      }
    }

    return  Response("Not Found", { status: 404 });
  },
});

console.log(`Listening on http://${ser.hostname}:${ser.port}...`);
