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
import * as utils$0 from "../utils/models.js";

// eslint-disable-next-line @typescript-eslint/ban-ts-comment
// @ts-ignore: Unused imports
import * as $models from "./models.js";

export function Count(): Promise<number> & { cancel(): void } {
    let $resultPromise = $Call.ByID(2521206865) as any;
    return $resultPromise;
}

export function Create(item: $models.WorkResult): Promise<$models.WorkResult> & { cancel(): void } {
    let $resultPromise = $Call.ByID(509814126, item) as any;
    let $typingPromise = $resultPromise.then(($result: any) => {
        return $$createType0($result);
    }) as any;
    $typingPromise.cancel = $resultPromise.cancel.bind($resultPromise);
    return $typingPromise;
}

export function Delete(id: number): Promise<void> & { cancel(): void } {
    let $resultPromise = $Call.ByID(2727636165, id) as any;
    return $resultPromise;
}

export function GetAll(): Promise<($models.WorkResult | null)[]> & { cancel(): void } {
    let $resultPromise = $Call.ByID(175211989) as any;
    let $typingPromise = $resultPromise.then(($result: any) => {
        return $$createType2($result);
    }) as any;
    $typingPromise.cancel = $resultPromise.cancel.bind($resultPromise);
    return $typingPromise;
}

export function GetById(id: number): Promise<$models.WorkResult | null> & { cancel(): void } {
    let $resultPromise = $Call.ByID(3823744434, id) as any;
    let $typingPromise = $resultPromise.then(($result: any) => {
        return $$createType1($result);
    }) as any;
    $typingPromise.cancel = $resultPromise.cancel.bind($resultPromise);
    return $typingPromise;
}

export function GetCompMaterialSpecs(): Promise<($models.WoodSpec | null)[]> & { cancel(): void } {
    let $resultPromise = $Call.ByID(1587952728) as any;
    let $typingPromise = $resultPromise.then(($result: any) => {
        return $$createType5($result);
    }) as any;
    $typingPromise.cancel = $resultPromise.cancel.bind($resultPromise);
    return $typingPromise;
}

export function GetCompProductSpecs(): Promise<($models.WoodSpec | null)[]> & { cancel(): void } {
    let $resultPromise = $Call.ByID(3029514276) as any;
    let $typingPromise = $resultPromise.then(($result: any) => {
        return $$createType5($result);
    }) as any;
    $typingPromise.cancel = $resultPromise.cancel.bind($resultPromise);
    return $typingPromise;
}

export function SearchByAllTextFields(phrase: string): Promise<($models.WorkResult | null)[]> & { cancel(): void } {
    let $resultPromise = $Call.ByID(2356462020, phrase) as any;
    let $typingPromise = $resultPromise.then(($result: any) => {
        return $$createType2($result);
    }) as any;
    $typingPromise.cancel = $resultPromise.cancel.bind($resultPromise);
    return $typingPromise;
}

export function SortedByOrder(fieldsSortingOrder: utils$0.SortField[]): Promise<($models.WorkResult | null)[]> & { cancel(): void } {
    let $resultPromise = $Call.ByID(3262837822, fieldsSortingOrder) as any;
    let $typingPromise = $resultPromise.then(($result: any) => {
        return $$createType2($result);
    }) as any;
    $typingPromise.cancel = $resultPromise.cancel.bind($resultPromise);
    return $typingPromise;
}

export function Update(item: $models.WorkResult): Promise<$models.WorkResult> & { cancel(): void } {
    let $resultPromise = $Call.ByID(805686227, item) as any;
    let $typingPromise = $resultPromise.then(($result: any) => {
        return $$createType0($result);
    }) as any;
    $typingPromise.cancel = $resultPromise.cancel.bind($resultPromise);
    return $typingPromise;
}

// Private type creation functions
const $$createType0 = models$0.WorkResult.createFrom;
const $$createType1 = $Create.Nullable($$createType0);
const $$createType2 = $Create.Array($$createType1);
const $$createType3 = models$0.WoodSpec.createFrom;
const $$createType4 = $Create.Nullable($$createType3);
const $$createType5 = $Create.Array($$createType4);
