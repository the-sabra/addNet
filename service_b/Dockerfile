# Use the official Bun image as a base image
FROM oven/bun:latest

# Set the working directory
WORKDIR /app

# Copy the package files
COPY package.json bun.lock ./

# Install dependencies
RUN bun install

# Copy the source code
COPY . .

# Expose the HTTP port
EXPOSE 2113

# Command to run the application
CMD ["bun", "run", "index.ts"]