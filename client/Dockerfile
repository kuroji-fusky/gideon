FROM node:lts AS base

ENV PNPM_HOME="/pnpm"
ENV PATH="$PNPM_HOME:$PATH"
RUN corepack enable

WORKDIR /client

FROM base as prod-dependencies
COPY package.json .
RUN --mount=type=cache,id=pnpm,target=/pnpm/store pnpm install --prod --frozen-lockfile

FROM base AS build
COPY . /client/.
RUN pnpm run build

FROM base
COPY --from=prod-dependencies ./client/node_modules /node_modules
COPY --from=build ./client/dist /dist

ENV HOST=0.0.0.0
ENV PORT=4321
EXPOSE 4321

CMD node ./dist/server/entry.mjs