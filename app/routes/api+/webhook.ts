import type { ActionFunctionArgs } from '@remix-run/node'


export const ROUTE_PATH = '/api/webhook' as const


export async function action({ request }: ActionFunctionArgs) {
  return new Response(null)
}
