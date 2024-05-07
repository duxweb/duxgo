import { useResource, useCustom, useParsed, useTranslate, useList } from '@refinedev/core'
import { CascaderAsync, FormPage } from '@duxweb/dux-refine'
import { MagicFormRender } from '@duxweb/dux-extend'
import { Form, Cascader } from 'tdesign-react/esm'

const Page = () => {
  const { id, params } = useParsed()
  const { resource } = useResource()
  const translate = useTranslate()
  const magic = params?.name

  const { data } = useCustom<Record<string, any>>({
    url: 'tools/magic/config',
    method: 'get',
    meta: {
      params: {
        magic: magic,
      },
    },
  })

  if (resource?.meta) {
    resource.meta.label = data?.data?.label
  }

  return (
    <FormPage
      formProps={{
        labelAlign: 'top',
      }}
      back
      id={id}
      queryParams={{
        magic: magic,
      }}
      useFormProps={{
        queryOptions: {
          cacheTime: 0,
        },
      }}
    >
      {data?.data?.type == 'tree' && (
        <Form.FormItem label={translate('tools.data.fields.parent')} name='parent_id'>
          <CascaderAsync
            url='tools/data'
            checkStrictly
            query={{
              magic: magic,
            }}
            keys={{
              label: data?.data?.tree_label || 'name',
              value: 'id',
            }}
          />
        </Form.FormItem>
      )}
      {data?.data?.fields && <MagicFormRender fields={data?.data?.fields} />}
    </FormPage>
  )
}

export default Page
