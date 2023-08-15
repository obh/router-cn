'use client';

import { Breadcrumb } from 'flowbite-react';
import { HiHome } from 'react-icons/hi';

export default function DefaultBreadcrumb({value}) {
  return (
    <Breadcrumb aria-label="Default breadcrumb example" className="py-8 px-4">
      <Breadcrumb.Item
        href="#"
        icon={HiHome}
      >
        <p>
          Home
        </p>
      </Breadcrumb.Item>
      <Breadcrumb.Item href="#">
        Payments
      </Breadcrumb.Item>
      <Breadcrumb.Item>
        {value}
      </Breadcrumb.Item>
    </Breadcrumb>
  )
}


