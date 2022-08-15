import React,{ useState,useMemo } from "react";
import useRestProps from '../hooks/useRestProps';
import useViewStyles from '../hooks/useViewStyles';

interface ViewProps extends Omit<React.HTMLProps<HTMLInputElement>, 'style'|'type'> {
    style?: React.CSSProperties | React.CSSProperties[];
    type?: "text" | "number" | "email" | "password" | any;
    initialValue?: string;
    multiple?: boolean;
    resize?: 'none' | 'both' | 'horizontal' | 'vertical';
    onChangeText?: (value: string) => void;
}

const TextInputer = (props:ViewProps) => {
    const type = props.type ?? "text";
    const [value,setValue] = useState(props.initialValue ?? '')

    const restProps = useRestProps(props,['children', 'style','className']);
    const viewStyles = useViewStyles(props.style,[{resize: props.resize ?? 'none'}])
    const className = useMemo(() => 'col-fs-fs'+(props.className ?? ''), [props.className])

    const onChange = (v:any) => {
        if(v.nativeEvent.inputType.indexOf('delete') !== -1){
            //delete char
            if(value.length === 0) return;
            let v = value.substring(0, value.length - 1)
            setValue(v);
            props.onChangeText && props.onChangeText(v);
        }else{
            //append char
            let nv = value + v.nativeEvent.data ?? ''
            setValue(nv);
            props.onChangeText && props.onChangeText(v);
        }
    }
    if(props.multiple) return (
        <textarea className={className} name="text" type={type} value={value} onChange={onChange} style={viewStyles} {...restProps}/>
    )
    return <input className={className} name="text" type={type} value={value} onChange={onChange} style={viewStyles} {...restProps}/>
}

export default TextInputer;