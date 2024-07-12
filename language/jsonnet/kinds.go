// Copyright 2019 VMware, Inc.
// SPDX-License-Identifier: Apache-2.0

package jsonnet

import "github.com/bazelbuild/bazel-gazelle/rule"

const (
	libraryRule       = "jsonnet_library"
	libraryRulePrefix = "library"

	toJSONRule       = "jsonnet_to_json"
	toJSONRulePrefix = "to_json"
)

var (
	// https://github.com/bazelbuild/rules_jsonnet
	jsonnetKinds = map[string]rule.KindInfo{
		libraryRule: {
			NonEmptyAttrs:  map[string]bool{"srcs": true},
			MergeableAttrs: map[string]bool{"srcs": true},
			ResolveAttrs:   map[string]bool{"deps": true},
		},
		toJSONRule: {
			NonEmptyAttrs: map[string]bool{
				"src":  true,
				"outs": true,
			},
			MergeableAttrs: map[string]bool{
				"src":  true,
				"outs": true,
			},
			ResolveAttrs: map[string]bool{"deps": true},
		},
	}
	jsonnetLoads = []rule.LoadInfo{
		{
			Name: "@rules_jsonnet//jsonnet:jsonnet.bzl",
			Symbols: []string{
				libraryRule,
				toJSONRule,
			},
		},
	}
)

func (*Lang) Kinds() map[string]rule.KindInfo { return jsonnetKinds }
func (*Lang) Loads() []rule.LoadInfo          { return jsonnetLoads }
