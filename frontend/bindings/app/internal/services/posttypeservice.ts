// Cynhyrchwyd y ffeil hon yn awtomatig. PEIDIWCH Â MODIWL
// This file is automatically generated. DO NOT EDIT

// eslint-disable-next-line @typescript-eslint/ban-ts-comment
// @ts-ignore: Unused imports
import {Call as $Call, Create as $Create} from "@wailsio/runtime";

// eslint-disable-next-line @typescript-eslint/ban-ts-comment
// @ts-ignore: Unused imports
import * as models$0 from "../models/models.js";

// eslint-disable-next-line @typescript-eslint/ban-ts-comment
// @ts-ignore: Unused imports
import * as $models from "./models.js";

export function Count(): Promise<number> & { cancel(): void } {
    let $resultPromise = $Call.ByID(554487955) as any;
    return $resultPromise;
}

export function Create(item: $models.PostType): Promise<$models.PostType> & { cancel(): void } {
    let $resultPromise = $Call.ByID(1092898136, item) as any;
    let $typingPromise = $resultPromise.then(($result: any) => {
        return $$createType0($result);
    }) as any;
    $typingPromise.cancel = $resultPromise.cancel.bind($resultPromise);
    return $typingPromise;
}

export function Delete(id: number): Promise<void> & { cancel(): void } {
    let $resultPromise = $Call.ByID(2114913543, id) as any;
    return $resultPromise;
}

export function GetAll(): Promise<($models.PostType | null)[]> & { cancel(): void } {
    let $resultPromise = $Call.ByID(416231387) as any;
    let $typingPromise = $resultPromise.then(($result: any) => {
        return $$createType2($result);
    }) as any;
    $typingPromise.cancel = $resultPromise.cancel.bind($resultPromise);
    return $typingPromise;
}

export function GetById(id: number): Promise<$models.PostType | null> & { cancel(): void } {
    let $resultPromise = $Call.ByID(3237123344, id) as any;
    let $typingPromise = $resultPromise.then(($result: any) => {
        return $$createType1($result);
    }) as any;
    $typingPromise.cancel = $resultPromise.cancel.bind($resultPromise);
    return $typingPromise;
}

export function Update(item: $models.PostType): Promise<$models.PostType> & { cancel(): void } {
    let $resultPromise = $Call.ByID(2773888269, item) as any;
    let $typingPromise = $resultPromise.then(($result: any) => {
        return $$createType0($result);
    }) as any;
    $typingPromise.cancel = $resultPromise.cancel.bind($resultPromise);
    return $typingPromise;
}

// Private type creation functions
const $$createType0 = models$0.PostType.createFrom;
const $$createType1 = $Create.Nullable($$createType0);
const $$createType2 = $Create.Array($$createType1);
