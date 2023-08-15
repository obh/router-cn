import { useRouter } from 'next/router'
import prisma from '../../lib/prisma'

import Head from 'next/head'
import { Table, Badge } from 'flowbite-react';
import DefaultSidebar from '../../components/DefaultSidebar';
import DefaultBreadcrumb from '../../components/DefaultBreadCrumb';


export default function PaymentAttempt(props) {

  const router = useRouter()
  console.log("received props --> ", props)
  const paymentattempts = props.paymentattempts
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
        <DefaultBreadcrumb value="Payment Attempt"/>
        <Table aria-label="Example static collection table">
          <Table.Head>
            <Table.HeadCell className="py-2 px-4">Id</Table.HeadCell>
            <Table.HeadCell className="py-2 px-4">Payment Intent</Table.HeadCell>
            <Table.HeadCell className="py-2 px-4">Amount</Table.HeadCell>
            <Table.HeadCell className="py-2 px-4">Currency</Table.HeadCell>
            <Table.HeadCell className="py-2 px-4">Provider</Table.HeadCell>
            <Table.HeadCell className="py-2 px-4">Status</Table.HeadCell>
            <Table.HeadCell className="py-2 px-4">Token</Table.HeadCell>
            <Table.HeadCell className="py-2 px-4">Added On</Table.HeadCell>
          </Table.Head>
          <Table.Body className="divide-y">
            {paymentattempts.map((pa) => (            
              <Table.Row className="bg-white dark:border-gray-700 dark:bg-gray-800">
                <Table.Cell>{pa.id}</Table.Cell>
                <Table.Cell>{pa.payment_intent_id}</Table.Cell>
                <Table.Cell>{pa.amount}</Table.Cell>
                <Table.Cell>{pa.currency}</Table.Cell>
                <Table.Cell>{pa.processor}</Table.Cell>
                <Table.Cell>{getBadge(pa.status)}</Table.Cell>
                <Table.Cell>{pa.payment_hash}</Table.Cell>
                <Table.Cell>{pa.added_on}</Table.Cell>
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
// export async function getStaticPaths(){
//   return {
//     paths: [
//       {
//         params: {  paymentSlug: '1',},
//       }, 
//       {
//         params: {  paymentSlug: '2',},
//       },
//     ],
//     fallback: false, // false or "blocking"
//   }
// }
  
export async function getServerSideProps({params}) {
  const paymentIntentId = parseInt(params.paymentSlug)
  const data = await prisma.PaymentAttempt.findMany({
    where: {
      "payment_intent_id": paymentIntentId
    }
  })
  //convert decimal value to string to pass through as json
  const paymentattempts = data.map((pa) => ({
    ...pa,
    amount: pa.amount.toString(),
    added_on: Math.floor(pa.added_on / 1000),
    updated_on: Math.floor(pa.updated_on / 1000)
  }))
  return {
    props: { paymentattempts },
  }
}

function getBadge(status){
  if(status == 'CREATED') {
    return (<Badge color="gray">{status}</Badge>)
  } else if(status == 'INITIATED'){
    return (<Badge color="warning">{status}</Badge>)
  } else if(status == 'REDIRECTED'){
    return (<Badge color="indigo">{status}</Badge>)
  }
}
  