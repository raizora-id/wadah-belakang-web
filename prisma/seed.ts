
import { PrismaClient } from '@prisma/client'
import { prisma } from '#app/utils/db.server'

const client = new PrismaClient()

async function seed() {
  /**
   * Users, Roles and Permissions.
   */
  const entities = ['user']
  const actions = ['create', 'read', 'update', 'delete']
  const accesses = ['own', 'any'] as const
  for (const entity of entities) {
    for (const action of actions) {
      for (const access of accesses) {
        await prisma.permission.create({ data: { entity, action, access } })
      }
    }
  }

  await prisma.role.create({
    data: {
      name: 'admin',
      permissions: {
        connect: await prisma.permission.findMany({
          select: { id: true },
          where: { access: 'any' },
        }),
      },
    },
  })
  await prisma.role.create({
    data: {
      name: 'user',
      permissions: {
        connect: await prisma.permission.findMany({
          select: { id: true },
          where: { access: 'own' },
        }),
      },
    },
  })
  await prisma.user.create({
    select: { id: true },
    data: {
      email: 'admin@admin.com',
      username: 'admin',
      roles: { connect: [{ name: 'admin' }, { name: 'user' }] },
    },
  })

  console.info(`🎭 User roles and permissions has been successfully created.`)

  console.info(`👒 Stripe Customer Portal has been successfully configured.`)
  console.info(
    '🎉 Visit: https://dashboard.stripe.com/test/products to see your products.',
  )
}

seed()
  .catch((err: unknown) => {
    console.error(err)
    process.exit(1)
  })
  .finally(async () => {
    await client.$disconnect()
  })
