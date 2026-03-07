import { readFile } from "node:fs/promises";
import { fileURLToPath } from "node:url";
import matter from "gray-matter";
import { marked } from "marked";

const ROOT_DIRECTORY = fileURLToPath(new URL("../../", import.meta.url));

export async function loadMarkdownPage(fileName) {
  const fullPath = `${ROOT_DIRECTORY}${fileName}`;
  const fileContent = await readFile(fullPath, "utf8");
  const { data, content } = matter(fileContent);

  return {
    frontmatter: data,
    content,
    html: marked.parse(content)
  };
}
