import React from "react";
import {useState} from "react";
import {useForm} from "react-hook-form";

interface FormState {
    source: string;
}
export const PostEditor = () => {
    const {
        register,
        handleSubmit,
        watch,
        formState: { errors }
    } = useForm<FormState>();

    const source = watch('source');
    const onSubmit = () => {};

    return (
        <form onSubmit={handleSubmit(onSubmit)}>
            <div id="post-form">
                <div className="flex-row">
                    <textarea id='source-textarea' {...register('source')} className="flex-row__fill"/>
                    <iframe id='preview' srcDoc={source} className="flex-row__fill"/>
                </div>
                <div className="flex-row">
                    <div className="flex-row__spacer" />
                    <button type="submit" className="flex-row__right">Save</button>
                </div>
            </div>
        </form>
    );
}