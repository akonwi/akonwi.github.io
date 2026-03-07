import rss from "@astrojs/rss";
import { getPosts } from "../lib/posts";
import { site } from "../site";

export async function GET(context) {
  const posts = await getPosts();

  return rss({
    title: site.title,
    description: site.description,
    site: context.site,
    items: posts.map((post) => ({
      title: post.title,
      pubDate: post.date,
      description: post.excerpt,
      link: post.url
    }))
  });
}
