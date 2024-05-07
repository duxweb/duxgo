import { useTranslate } from '@refinedev/core'
import { CodeEditor, FormModal } from '@duxweb/dux-refine'
import { Form, Input, Radio } from 'tdesign-react/esm'

const Page = (props: Record<string, any>) => {
  const translate = useTranslate()

  return (
    <FormModal id={props?.id}>
      <Form.FormItem label={translate('tools.magicSource.fields.name')} name='name'>
        <Input />
      </Form.FormItem>
      <Form.FormItem label={translate('tools.magicSource.fields.type')} name='type' initialData={0}>
        <Radio.Group>
          <Radio value={0}>{translate('tools.magicSource.fields.data')}</Radio>
          <Radio value={1}>{translate('tools.magicSource.fields.async')}</Radio>
        </Radio.Group>
      </Form.FormItem>
      <Form.FormItem shouldUpdate={(prev, next) => prev.type !== next.type}>
        {({ getFieldValue }) => {
          if (getFieldValue('type') === 0) {
            return (
              <Form.FormItem
                label={translate('tools.magicSource.fields.data')}
                name='data'
                key='data'
              >
                <CodeEditor type='json' />
              </Form.FormItem>
            )
          }
          if (getFieldValue('type') === 1) {
            return (
              <>
                <Form.FormItem
                  label={translate('tools.magicSource.fields.url')}
                  name='url'
                  key={'url'}
                >
                  <Input />
                </Form.FormItem>
              </>
            )
          }
          return <></>
        }}
      </Form.FormItem>
    </FormModal>
  )
}

export default Page
