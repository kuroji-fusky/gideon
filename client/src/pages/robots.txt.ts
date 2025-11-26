import type { APIRoute } from "astro"

const robotsTxt = `
# Ensures that wiki mirrors won't get indexed by most search engines
# If you change this, that's on you
User-agent: *
Disallow:

Sitemap: ${new URL("sitemap-index.xml", import.meta.env.SITE).href}
`

export const GET: APIRoute = () => {
  return new Response(robotsTxt.trim(), {
    headers: {
      "Content-Type": "text/plain; charset=utf-8"
    }
  })
}