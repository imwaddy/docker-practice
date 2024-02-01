# Use an official Node.js image as a base image
FROM node:14 AS builder

# Set the working directory
WORKDIR /app

# Copy package.json and package-lock.json to the container
COPY package*.json ./

# Install dependencies
RUN npm install

# Copy the rest of the application code
COPY . .

# Build the Vue.js application
RUN npm run build

# Use a lightweight Nginx image as the final image
FROM nginx:alpine

# Copy the built files from the builder stage to the Nginx server
COPY --from=builder /app/dist /usr/share/nginx/html

# Expose port 8081
EXPOSE 8081

# Command to start Nginx
CMD ["nginx", "-g", "daemon off;"]
