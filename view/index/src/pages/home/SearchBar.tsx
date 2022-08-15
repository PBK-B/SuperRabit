import { Input } from 'antd';
import { CloseCircleFilled, PlusCircleOutlined } from '@ant-design/icons'
import StyleSheet from '@/components/StyleSheet';

const SearchBar = (props: {
    onAddClick?: () => void
}) => {
    const { onAddClick } = props
    return (
        <div style={styles.bar}>
            <div style={styles.bar_content}>
                <div style={styles.bar_content_input}>
                    <Input
                        type="text"
                        bordered={false}
                        placeholder="搜索你想下载的App"
                        style={styles.input}
                    />
                    {/* <CloseCircleFilled style={styles.close} /> */}
                </div>
                {/* <PlusCircleOutlined onClick={onAddClick} style={{
                    fontSize: 21,
                    color: '#666',
                    marginRight: 15,
                }} /> */}
            </div>
        </div>
    )
}
export default SearchBar

const styles = StyleSheet.create({
    bar: {
        width: '100vw',
        height: 56,
        background: 'white',
        // boxShadow: '0 0px 8px #00000022',
        position: 'fixed',
        top: 0,
        zIndex: 1,
    },
    bar_content: {
        display: 'flex',
        width: '100%',
        height: '100%',
        justifyContent: 'center',
        alignItems: 'center'
    },
    bar_content_input: {
        display: 'flex',
        flexGrow: 1,
        margin: '0 15px',
        height: 32,
        alignItems: 'center',
        borderRadius: 8,
        background: '#f8f8fa'
    },
    input: {
        display: 'flex',
        border: 'none',
        backgroundColor: 'transparent',
        fontSize: 14,
    },
    close: {
        display: 'flex',
        margin: '0 12px',
        color: '#bbb',
    }
})