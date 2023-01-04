import React from 'react'
import '../styles/SideBarItem.css'
import { Typography } from '@mui/material'
import { AppPageName } from '../enums/AppPageName'

export type SideBarItemProps = {
    name: string,
    setCurrentPageName: React.Dispatch<React.SetStateAction<AppPageName>>,
    isCurrent: boolean
}

const SideBarItem = (props: SideBarItemProps) => {

    const onClick = () => {
        switch (props.name) {
            case AppPageName.BALANCE:
                props.setCurrentPageName(AppPageName.BALANCE);
                break;
            case AppPageName.TRANSACTIONS:
                props.setCurrentPageName(AppPageName.TRANSACTIONS);
                break;
            case AppPageName.TRANSFER:
                props.setCurrentPageName(AppPageName.TRANSFER);
                break;
            default:
                props.setCurrentPageName(AppPageName.DASHBOARD);
        }
    }

  return (
    <div onClick={onClick} className={`sidebaritem ${props.isCurrent ? "sidebaritem-selected" : ""}`}>
        <Typography>{props.name}</Typography>
    </div>
  )
}

export default SideBarItem