import { authMiddleware } from '@clerk/nextjs';

export default authMiddleware()

// protected routes
export const config = {
    matcher: ["/dashboard", "/profile"]
}