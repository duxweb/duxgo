import React, { useMemo, useRef, useState } from 'react'
import { useTranslate, useList, useDeleteMany } from '@refinedev/core'
import {
  PrimaryTableCol,
  Link,
  Select,
  Form,
  Button,
  Tree,
  List as TdList,
  Space,
  Menu,
  Divider,
  Dropdown,
  MenuValue,
  DialogPlugin,
} from 'tdesign-react/esm'
import {
  PageTable,
  useSelect,
  TableRef,
  ButtonModal,
  DeleteButton,
  DeleteLink,
  EmptyWidget,
  CardSider, Modal,
} from '@duxweb/dux-refine'
import clsx from 'clsx'

interface FileIconProps {
  mime: string
}
const FileIcon = ({ mime }: FileIconProps) => {
  switch (true) {
    case /^image\//.test(mime):
      return (
        <div className='h-10 w-10 flex items-center justify-center rounded p-2 text-white bg-brand'>
          <div className='i-tabler:photo h-6 w-6'></div>
        </div>
      )
    case /^video\//.test(mime):
      return (
        <div className='h-10 w-10 flex items-center justify-center rounded p-2 text-white bg-success'>
          <div className='i-tabler:video h-6 w-6'></div>
        </div>
      )
    case /^audio\//.test(mime):
      return (
        <div className='h-10 w-10 flex items-center justify-center rounded p-2 text-white bg-warning'>
          <div className='i-tabler:audio h-6 w-6'></div>
        </div>
      )
    case /^application\/pdf$/.test(mime):
      return (
        <div className='h-10 w-10 flex items-center justify-center rounded p-2 text-white bg-error'>
          <div className='i-tabler:file-pdf h-6 w-6'></div>
        </div>
      )
    case /^application\/vnd\.openxmlformats-officedocument\.wordprocessingml\.document$/.test(mime):
    case /^application\/msword$/.test(mime):
      return (
        <div className='h-10 w-10 flex items-center justify-center rounded p-2 text-white bg-brand'>
          <div className='i-tabler:file-word h-6 w-6'></div>
        </div>
      )
    case /^application\/vnd\.openxmlformats-officedocument\.spreadsheetml\.sheet$/.test(mime):
    case /^application\/vnd\.ms-excel$/.test(mime):
      return (
        <div className='h-10 w-10 flex items-center justify-center rounded p-2 text-white bg-brand'>
          <div className='i-tabler:file-excel h-6 w-6'></div>
        </div>
      )
    case /^application\/zip$/.test(mime):
    case /^application\/x-rar-compressed$/.test(mime):
    case /^application\/x-7z-compressed$/.test(mime):
      return (
        <div className='h-10 w-10 flex items-center justify-center rounded p-2 text-white bg-brand'>
          <div className='i-tabler:file-zip h-6 w-6'></div>
        </div>
      )
    default:
      return (
        <div className='h-10 w-10 flex items-center justify-center rounded p-2 text-white bg-brand'>
          <div className='i-tabler:file-unknown h-6 w-6'></div>
        </div>
      )
  }
}

const List = () => {
  const translate = useTranslate()
  const table = useRef<TableRef>(null)

  const columns = React.useMemo<PrimaryTableCol[]>(
    () => [
      {
        colKey: 'name',
        title: translate('tools.file.fields.name'),
        minWidth: 300,
        cell: ({ row }) => {
          return (
            <div className='flex items-center gap-2'>
              <div>
                <FileIcon mime={row.mime} />
              </div>
              <div className='flex flex-col'>
                <div>{row.name}</div>
                <div className='text-sm text-gray'>{row.mime}</div>
              </div>
            </div>
          )
        },
      },
      {
        colKey: 'dir_name',
        title: translate('tools.file.fields.dir'),
        width: 150,
      },
      {
        colKey: 'size',
        title: translate('tools.file.fields.size'),
        width: 150,
      },
      {
        colKey: 'driver',
        title: translate('tools.file.fields.driver'),
        width: 150,
      },
      {
        colKey: 'time',
        title: translate('tools.file.fields.time'),
        width: 200,
      },
      {
        colKey: 'link',
        title: translate('table.actions'),
        fixed: 'right',
        align: 'center',
        width: 160,
        cell: ({ row }) => {
          return (
            <div className='flex justify-center gap-4'>
              <Link theme='primary' href={row.url} target='_block'>
                {translate('buttons.show')}
              </Link>
              <DeleteLink rowId={row.id} />
            </div>
          )
        },
      },
    ],
    // eslint-disable-next-line react-hooks/exhaustive-deps
    [translate],
  )

  const dirId = table.current?.filters?.dir_id

  return (
    <>
      <PageTable
        ref={table}
        columns={columns}
        table={{
          rowKey: 'id',
        }}
        siderRender={() => <SideTree />}
        // filterRender={() => (
        //   <>
        //     <div>
        //       <ButtonModal
        //         resource='tools.fileDir'
        //         action='create'
        //         variant='outline'
        //         theme='default'
        //         title={translate('tools.file.fields.addDir')}
        //         icon={<div className='i-tabler:plus' />}
        //         component={() => import('./group')}
        //       >
        //         <></>
        //       </ButtonModal>
        //     </div>
        //     <Form.FormItem name='dir_id' initialData={options?.[0]?.value}>
        //       <Select
        //         filterable
        //         clearable
        //         onSearch={onSearch}
        //         options={options}
        //         placeholder={translate('tools.file.placeholder.dir')}
        //         loading={queryResult.isLoading}
        //       />
        //     </Form.FormItem>

        //     {table.current?.filters?.dir_id && (
        //       <div className='flex gap-2'>
        //         <ButtonModal
        //           resource='tools.fileDir'
        //           action='edit'
        //           variant='outline'
        //           theme='default'
        //           title={translate('tools.file.fields.editDir')}
        //           icon={<div className='i-tabler:edit' />}
        //           component={() => import('./group')}
        //           rowId={table.current?.filters.dir_id}
        //         >
        //           <></>
        //         </ButtonModal>
        //         <DeleteButton
        //           resource='tools.fileDir'
        //           variant='outline'
        //           theme='default'
        //           icon={<div className='i-tabler:trash' />}
        //           rowId={table.current?.filters.dir_id}
        //         >
        //           <></>
        //         </DeleteButton>
        //       </div>
        //     )}
        //   </>
        // )}
        actionRender={() => (
          <ButtonModal
            component={() => import('./upload')}
            action='upload'
            title={translate('tools.file.fields.upload')}
            icon={<div className='t-icon i-tabler:plus'></div>}
            rowId={dirId}
          />
        )}
      />
    </>
  )
}

const SideTree = () => {
  const translate = useTranslate()

  const { data, isLoading } = useList({
    resource: 'tools.fileDir',
  })


  const { mutate } = useDeleteMany()

  const [value, setValue] = useState<MenuValue>()

  return (
    <CardSider
      title='附件分类'
      tools={
        <>
          <ButtonModal
            resource='tools.fileDir'
            action='create'
            variant='text'
            shape='circle'
            theme='default'
            title={translate('tools.file.fields.addDir')}
            icon={<div className='i-tabler:plus' />}
            component={() => import('./group')}
          >
            <></>
          </ButtonModal>
          <Button
            action='create'
            variant='text'
            shape='circle'
            theme='default'
            icon={<div className='i-tabler:refresh' />}
            onClick={() => setValue('0')}
          />
          {value && value != '0' && (
            <Dropdown
              direction='right'
              hideAfterItemClick
              placement='bottom-left'
              trigger='hover'
            >
              <Button
                action='create'
                variant='text'
                shape='circle'
                theme='default'
                icon={<div className='i-tabler:dots-vertical' />}
              />
              <Dropdown.DropdownMenu>
                <Dropdown.DropdownItem value={1} prefixIcon={<div className='i-tabler:edit'></div>} onClick={() => {
                  Modal.open({
                    title: '修改目录',
                    component: () => import('./group'),
                    componentProps: {
                      id: value
                    }
                  })
                }}>编辑</Dropdown.DropdownItem>
                <Dropdown.DropdownItem value={2} prefixIcon={<div className='i-tabler:x'></div>} onClick={() => {
                  const confirmDia = DialogPlugin.confirm({
                    className: 'app-modal',
                    header: '确认删除',
                    width: 350,
                    body: <div className='p-4'>确认执行该操作？</div>,
                    onClose: ({ e, trigger }) => {
                      console.log('e: ', e);
                      console.log('trigger: ', trigger);
                      confirmDia.hide();
                    },
                    onConfirm: () => {
                      mutate({
                        resource: 'tools.fileDir',
                        ids: [value],
                      })
                      confirmDia.hide();
                    }
                  })
                }}>删除</Dropdown.DropdownItem>
              </Dropdown.DropdownMenu>
            </Dropdown>
          )}
        </>
      }
    >
      {data?.data && data?.data.length > 0 ? <Menu
        expandType='normal'
        theme='light'
        width={'100%'}
        className='app-sider-menu bg-transparent'
        value={value}
        onChange={(v) => {
          setValue(v)
        }}
      >
        <SideTreeChilren data={data?.data} optionLabel='name' optionValue='id' />
      </Menu> : <div className='text-sm mt-4'><EmptyWidget type='simple' /></div>}
    </CardSider>
  )
}

interface SideTreeChilrenProps {
  data?: Record<string, any>[]
  optionLabel?: string
  optionValue?: string
  optionChildren?: string
}

const SideTreeChilren = ({
  data,
  optionLabel = 'label',
  optionValue = 'value',
  optionChildren = 'children',
}: SideTreeChilrenProps) => {
  return data?.map((item, k) => {
    if (item['optionChildren'] && item['optionChildren'].length > 0) {
      return (
        <Menu.SubMenu key={k} title={item[optionLabel]} value={item[optionValue]}>
          {SideTreeChilren({
            data: item[optionChildren] as Record<string, any>[],
          })}
        </Menu.SubMenu>
      )
    } else {
      return (
        <Menu.MenuItem key={k} value={item[optionValue]}>
          {item[optionLabel]}
        </Menu.MenuItem>
      )
    }
  })
}

export default List
