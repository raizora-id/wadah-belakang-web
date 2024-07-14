import type {
    MetaFunction,
    LoaderFunctionArgs,
    ActionFunctionArgs,
  } from '@remix-run/node'
  import { json } from '@remix-run/node'
  import { z } from 'zod'
  import { authenticator } from '#app/modules/auth/auth.server'
  import { getSession, commitSession } from '#app/modules/auth/auth-session.server'
  import { validateCSRF } from '#app/utils/csrf.server'
  import { checkHoneypot } from '#app/utils/honeypot.server'

  import { siteConfig } from '#app/utils/constants/brand'
  import { ROUTE_PATH as DASHBOARD_PATH } from '#app/routes/dashboard+/_layout'
  import { ROUTE_PATH as AUTH_VERIFY_PATH } from '#app/routes/auth+/verify'
  
  export const ROUTE_PATH = '/auth/login' as const
  
  export const LoginSchema = z.object({
    email: z.string().max(256).email('Email address is not valid.'),
  })
  
  export const meta: MetaFunction = () => {
    return [{ title: `${siteConfig.siteTitle} - Login` }]
  }
  
  export async function loader({ request }: LoaderFunctionArgs) {
    await authenticator.isAuthenticated(request, {
      successRedirect: DASHBOARD_PATH,
    })
  
    const cookie = await getSession(request.headers.get('Cookie'))
    const authEmail = cookie.get('auth:email')
    const authError = cookie.get(authenticator.sessionErrorKey)
  
    return json({ authEmail, authError } as const, {
      headers: {
        'Set-Cookie': await commitSession(cookie),
      },
    })
  }
  
  export async function action({ request }: ActionFunctionArgs) {
    const url = new URL(request.url)
    const pathname = url.pathname
  
    const clonedRequest = request.clone()
    const formData = await clonedRequest.formData()
    await validateCSRF(formData, clonedRequest.headers)
    checkHoneypot(formData)
  
    await authenticator.authenticate('TOTP', request, {
      successRedirect: AUTH_VERIFY_PATH,
      failureRedirect: pathname,
    })
  }
  
  export default function DetailProductPage() {
    return (
      <div className="mx-auto flex h-full w-full max-w-96 flex-col items-center justify-center gap-6">
        <div className="mb-2 flex flex-col gap-2">
          <h3 className="text-center text-2xl font-medium text-primary">
            Product Name
          </h3>
          <p className="text-center text-base font-normal text-primary/60">
            Welcome back! Please log in to continue.
          </p>
        </div>
      </div>
    )
  }
  