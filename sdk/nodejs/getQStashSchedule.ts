// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

import * as pulumi from "@pulumi/pulumi";
import * as utilities from "./utilities";

/**
 * ## Example Usage
 *
 * ```typescript
 * import * as pulumi from "@pulumi/pulumi";
 * import * as upstash from "@pulumi/upstash";
 *
 * const exampleQstashScheduleData = upstash.getQStashSchedule({
 *     scheduleId: resource.upstash_qstash_schedule.exampleQstashSchedule.schedule_id,
 * });
 * ```
 */
export function getQStashSchedule(args: GetQStashScheduleArgs, opts?: pulumi.InvokeOptions): Promise<GetQStashScheduleResult> {
    if (!opts) {
        opts = {}
    }

    opts = pulumi.mergeOptions(utilities.resourceOptsDefaults(), opts);
    return pulumi.runtime.invoke("upstash:index/getQStashSchedule:getQStashSchedule", {
        "scheduleId": args.scheduleId,
    }, opts);
}

/**
 * A collection of arguments for invoking getQStashSchedule.
 */
export interface GetQStashScheduleArgs {
    scheduleId: string;
}

/**
 * A collection of values returned by getQStashSchedule.
 */
export interface GetQStashScheduleResult {
    readonly body: string;
    readonly createdAt: number;
    readonly cron: string;
    readonly destination: string;
    readonly forwardHeaders: {[key: string]: string};
    /**
     * The provider-assigned unique ID for this managed resource.
     */
    readonly id: string;
    readonly notBefore: number;
    readonly retries: number;
    readonly scheduleId: string;
}

export function getQStashScheduleOutput(args: GetQStashScheduleOutputArgs, opts?: pulumi.InvokeOptions): pulumi.Output<GetQStashScheduleResult> {
    return pulumi.output(args).apply(a => getQStashSchedule(a, opts))
}

/**
 * A collection of arguments for invoking getQStashSchedule.
 */
export interface GetQStashScheduleOutputArgs {
    scheduleId: pulumi.Input<string>;
}