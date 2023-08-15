# Use an official Node.js image as the base image
FROM node:19.0.0-alpine3.16 as builder

# Set the working directory inside the container
WORKDIR /app

# Copy the package.json and package-lock.json files to the container
COPY package.json package-lock.json next.config.js /app/
COPY prisma /app/prisma/

# Install dependencies
RUN npm config ls -l
RUN npm install  

# Copy the rest of the application code to the container
COPY . /app/

#env
ARG DATABASE_URL
ENV DATABASE_URL $DATABASE_URL

# Build prisma
RUN npx prisma generate

# Build the Next.js application
RUN npm run build

FROM node:19.0.0-alpine3.16 as runner
ENV NODE_ENV=production
#removing node modules to reduce size of image (see standalone) (update: many things dont' work with standalone)

COPY --chown=node:node --from=builder /app/node_modules ./node_modules
COPY --chown=node:node --from=builder /app/package*.json ./
COPY --chown=node:node --from=builder /app/next.config.js ./
#COPY --chown=node:node --from=builder /app/.next/static ./.next/static
COPY --chown=node:node --from=builder /app/.next ./.next

# Expose the port that the application will run on
EXPOSE 3000

# Start the application
CMD ["npm", "start"]

