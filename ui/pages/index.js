import Head from 'next/head'
import Payment from '../components/Payment'
import prisma from '../lib/prisma'
import { Table, Badge } from 'flowbite-react';
import DefaultSidebar from '../components/DefaultSidebar';
import DefaultBreadcrumb from '../components/DefaultBreadCrumb';


export default function Home({ paymentintents }) {


  return (
    <div>
      <Head>
        <title>Cashfree Dashboard</title>
        <meta name="description" content="Cashfree Dashboard" />
        <link rel="icon" href="/favicon.ico" />
      </Head>
      <div className="flex items-start">
      <DefaultSidebar/>      
      
      <main className="relative h-full w-full overflow-y-auto bg-gray-10 dark:bg-gray-400">
        <DefaultBreadcrumb value="Payment Intent"/>
        <Table aria-label="Example static collection table">
          <Table.Head>
            <Table.HeadCell className="py-2 px-4">Id</Table.HeadCell>
            <Table.HeadCell className="py-2 px-4">Customer Email</Table.HeadCell>
            <Table.HeadCell className="py-2 px-4">Customer Phone</Table.HeadCell>
            <Table.HeadCell className="py-2 px-4">Amount</Table.HeadCell>
            <Table.HeadCell className="py-2 px-4">Currency</Table.HeadCell>
            <Table.HeadCell className="py-2 px-4">AddedOn</Table.HeadCell>
            <Table.HeadCell className="py-2 px-4">Status</Table.HeadCell>
            <Table.HeadCell className="py-2 px-4">Details</Table.HeadCell>
          </Table.Head>
          <Table.Body className="divide-y">
            {paymentintents.map((pi) => (            
              <Table.Row className="bg-white dark:border-gray-700 dark:bg-gray-800">
                <Table.Cell>{pi.id}</Table.Cell>
                <Table.Cell>{pi.customer_email}</Table.Cell>
                <Table.Cell>{pi.customer_phone}</Table.Cell>
                <Table.Cell>{pi.amount}</Table.Cell>
                <Table.Cell>{pi.currency}</Table.Cell>
                <Table.Cell>{pi.added_on}</Table.Cell>
                <Table.Cell><Badge color="info">{pi.status}</Badge></Table.Cell>
                <Table.Cell><a href='#'>link</a></Table.Cell>
              </Table.Row>
              // <Payment paymentIntent={pi} key={pi.id} />
            ))}
          </Table.Body>
        </Table>
      </main>
      </div>

      <footer>         
      </footer>
    </div>
  )
}


export async function getServerSideProps({context}) {
  const data = await prisma.paymentIntent.findMany({
  })

  //convert decimal value to string to pass through as json
  const paymentintents = data.map((pi) => ({
    ...pi,
    amount: pi.amount.toString(),
    added_on: Math.floor(pi.added_on / 1000),
    updated_on: Math.floor(pi.updated_on / 1000)
  }))
  return {
    props: { paymentintents },
  }
}
