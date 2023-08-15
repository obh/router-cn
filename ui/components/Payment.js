

export default function Payment({ paymentIntent }) {
  const { id, customer_email, customer_phone, amount, currency, added_on } = paymentIntent

  return (
    <div>
    <tr>
      <td>hello</td>
      
    </tr>
    </div>
  )
}

 // return (
  //   <div
  //     className="max-w-[250px] rounded overflow-hidden shadow-lg"
  //     key={paymentIntent.id}
  //   >
  //     <Image
  //       className="w-full"
  //       width={250}
  //       height={250}
  //       objectFit="cover"
  //       src={'/helmet.jpg'}
  //     />
  //     <div className="px-6 py-4">
  //       <div className="font-bold text-xl mb-2">{id}</div>
  //       <p className="text-gray-700 text-base">{customer_email}</p>
  //       <p className="text-gray-900 text-xl">${amount}</p>
  //     </div>
  //     <div className="px-6 pt-4 pb-2">
  //       <span className="inline-block bg-gray-200 rounded-full px-3 py-1 text-sm font-semibold text-gray-700 mr-2 mb-2">
  //         {customer_phone}
  //       </span>
  //     </div>
  //   </div>
  // )
