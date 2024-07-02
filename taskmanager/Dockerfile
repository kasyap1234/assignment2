# Use Node.js LTS version as base image
FROM node:lts AS build

# Set working directory inside the container
WORKDIR /app

# Copy package.json and package-lock.json to work directory
COPY package.json package-lock.json ./

# Install dependencies
RUN npm install

# Copy all source code to work directory
COPY . .

# Build the SvelteKit application
RUN npm run build

# Production stage
FROM node:lts-alpine

# Set working directory inside the container
WORKDIR /app

# Copy built assets from previous stage
COPY --from=build /app/build /app/build

# Expose port 3000
EXPOSE 3000

# Command to run the application
CMD ["npm", "run", "start"]
