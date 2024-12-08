import {NextRequest, NextResponse} from "next/server";
import {config} from "@/config/config";

export function middleware(request: NextRequest) {
    if (request.nextUrl.pathname.startsWith('/admin')) {
        const authHeader = request.headers.get('authorization');
        if (authHeader == null) {
            return new Response('Unauthorized', {
                status: 401,
                headers: {
                    'WWW-Authenticate': 'Basic realm="Access to the admin area"',
                },
            });
        }

        const [username, password] = atob(authHeader.split(' ')[1]).split(':');
        if (username !== config.ADMIN_USER || password !== config.ADMIN_PASSWORD) {
            return new Response('Authentication Failed', {
                status: 401,
                headers: {
                    'WWW-Authenticate': 'Basic realm="Access to the admin area"',
                }
            });
        }

        return NextResponse.next();
    } else {
        return NextResponse.next();
    }
}