import React from 'react'
import '../styles/Sidebar.css'
import SideBarItem from './SideBarItem'
import { Typography } from '@mui/material'
import { AppPageName } from '../enums/AppPageName'

const sideBarItemNames = [
    "DashBoard",
    "Transactions",
    "Balance",
    "Transfer"
]

type SideBarProps = {
    currentPageName: AppPageName,
    setCurrentPageName: React.Dispatch<React.SetStateAction<AppPageName>>
}

const SideBar = (props: SideBarProps) => {
    return (
        <div className='sidebar'>
            <Typography variant='h5' margin="5px">App Name</Typography>
            {
                sideBarItemNames.map((itemName) => {
                    if(props.currentPageName == itemName) {
                        return <SideBarItem key={itemName} name={itemName} setCurrentPageName={props.setCurrentPageName} isCurrent={true} />
                    }
                    return <SideBarItem key={itemName} name={itemName} setCurrentPageName={props.setCurrentPageName} isCurrent={false} />
                })
            }
        </div>
    )
}

export default SideBar