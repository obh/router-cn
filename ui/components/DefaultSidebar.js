'use client';

import { Sidebar } from 'flowbite-react';
import { HiArrowSmRight, HiChartPie, HiInbox, HiShoppingBag, HiTable, HiUser, HiViewBoards } from 'react-icons/hi';

export default function DefaultSidebar() {
  return (
    <Sidebar aria-label="Default sidebar example">
      <div className="flex h-full flex-col justify-between py-2 h-screen sticky top-0">
      <Sidebar.Items>
        <Sidebar.ItemGroup>
          <Sidebar.Item href="#" icon={HiChartPie}>
            <p>
              Dashboard
            </p>
          </Sidebar.Item>
          <Sidebar.Item
            href="#"
            icon={HiViewBoards}
            labelColor="dark"
          >
            <p>
              Payments
            </p>
          </Sidebar.Item>
          <Sidebar.Item
            href="#"
            icon={HiInbox}
          >
            <p>
              Refunds
            </p>
          </Sidebar.Item>
          <Sidebar.Item
            href="#"
            icon={HiUser}
          >
            <p>
              Settlements
            </p>
          </Sidebar.Item>
          {/* <Sidebar.Item
            href="#"
            icon={HiShoppingBag}
          >
            <p>
              Products
            </p>
          </Sidebar.Item>
          <Sidebar.Item
            href="#"
            icon={HiArrowSmRight}
          >
            <p>
              Sign In
            </p>
          </Sidebar.Item>
          <Sidebar.Item
            href="#"
            icon={HiTable}
          >
            <p>
              Sign Up
            </p>
          </Sidebar.Item> */}
        </Sidebar.ItemGroup>
      </Sidebar.Items>
      </div>
    </Sidebar>
  )
}


