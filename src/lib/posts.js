import { readdir, readFile } from "node:fs/promises";
import path from "node:path";
import { fileURLToPath } from "node:url";
import matter from "gray-matter";
import { marked } from "marked";
import { markedHighlight } from "marked-highlight";
import hljs from "highlight.js/lib/common";

const POSTS_DIRECTORY = fileURLToPath(new URL("../../_posts", import.meta.url));
const POST_FILENAME_REGEX = /^(\d{4})-(\d{1,2})-(\d{1,2})-(.+)\.(md|markdown)$/;

marked.setOptions({
  gfm: true,
  breaks: false
});

marked.use(
  markedHighlight({
    langPrefix: "hljs language-",
    highlight(code, language) {
      if (language && hljs.getLanguage(language)) {
        return hljs.highlight(code, { language }).value;
      }

      return hljs.highlightAuto(code).value;
    }
  })
);

function normalizeCategories(value) {
  if (Array.isArray(value)) {
    return value.map((entry) => String(entry));
  }

  if (typeof value === "string" && value.trim().length > 0) {
    return [value.trim()];
  }

  return [];
}

function parseDate(value, fallback) {
  const parsed = new Date(value);

  if (Number.isNaN(parsed.getTime())) {
    return fallback;
  }

  return parsed;
}

function stripMarkdown(markdown) {
  return markdown
    .replace(/```[\s\S]*?```/g, "")
    .replace(/`([^`]+)`/g, "$1")
    .replace(/!\[[^\]]*\]\([^\)]*\)/g, "")
    .replace(/\[([^\]]+)\]\([^\)]*\)/g, "$1")
    .replace(/^#+\s+/gm, "")
    .replace(/[>*_~]/g, "")
    .replace(/\s+/g, " ")
    .trim();
}

function getExcerpt(content) {
  const candidate = content
    .split(/\n\s*\n/g)
    .map((block) => block.trim())
    .find((block) => block.length > 0 && !block.startsWith("#"));

  if (!candidate) {
    return "";
  }

  const plainText = stripMarkdown(candidate);
  if (plainText.length <= 220) {
    return plainText;
  }

  return `${plainText.slice(0, 217)}...`;
}

function getSlugFromFilename(filename) {
  const match = filename.match(POST_FILENAME_REGEX);
  if (!match) {
    return null;
  }

  return match[4];
}

function getDateFromFilename(filename) {
  const match = filename.match(POST_FILENAME_REGEX);
  if (!match) {
    return null;
  }

  const year = Number(match[1]);
  const month = Number(match[2]);
  const day = Number(match[3]);

  return new Date(Date.UTC(year, month - 1, day));
}

export function formatPostDate(dateValue) {
  return new Intl.DateTimeFormat("en-US", {
    month: "short",
    day: "numeric",
    year: "numeric"
  }).format(dateValue);
}

export async function getPosts() {
  const fileNames = await readdir(POSTS_DIRECTORY);
  const markdownFiles = fileNames.filter((name) => /\.(md|markdown)$/.test(name));

  const entries = await Promise.all(
    markdownFiles.map(async (fileName) => {
      const slug = getSlugFromFilename(fileName);
      if (!slug) {
        return null;
      }

      const fallbackDate = getDateFromFilename(fileName) ?? new Date(0);
      const fullPath = path.join(POSTS_DIRECTORY, fileName);
      const fileContent = await readFile(fullPath, "utf8");
      const { data, content } = matter(fileContent);
      const published = data.published !== false;

      if (!published) {
        return null;
      }

      const postDate = parseDate(data.date, fallbackDate);

      return {
        slug,
        fileName,
        title: data.title ? String(data.title) : slug,
        date: postDate,
        categories: normalizeCategories(data.categories),
        content,
        excerpt: getExcerpt(content),
        html: marked.parse(content),
        url: `/blog/${slug}`
      };
    })
  );

  return entries
    .filter((entry) => entry !== null)
    .sort((left, right) => right.date.getTime() - left.date.getTime());
}

export async function getPostBySlug(slug) {
  const posts = await getPosts();
  return posts.find((post) => post.slug === slug) ?? null;
}
