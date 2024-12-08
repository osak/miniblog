'use client';

import React from "react";
import {useForm} from "react-hook-form";

interface FormState {
    source: string;
}
export const PostEditor = () => {
    const {
        register,
        handleSubmit,
        watch,
    } = useForm<FormState>();

    const source = watch('source');
    const onSubmit = () => {};

    return (
        <form onSubmit={handleSubmit(onSubmit)}>
            <div id="post-form">
                <div className="flex">
                    <textarea id='source-textarea' {...register('source')} className="grow"/>
                    <iframe id='preview' srcDoc={source} className="grow"/>
                </div>
                <div className="flex">
                    <div className="grow" />
                    <button type="submit">Save</button>
                </div>
            </div>
        </form>
    );
}