package gotests

import (
	"errors"
	"go/types"
	"io/ioutil"
	"path"
	"regexp"
	"testing"
	"unicode"
)

func TestGenerateTests(t *testing.T) {
	tests := []struct {
		name              string
		srcPath           string
		only              *regexp.Regexp
		excl              *regexp.Regexp
		exported          bool
		printInputs       bool
		importer          types.Importer
		want              string
		wantNoTests       bool
		wantMultipleTests bool
		wantErr           bool
	}{
		{
			name:        "Blank Go file",
			srcPath:     `testdata/blankfile/blank.go`,
			wantNoTests: true,
			wantErr:     true,
		}, {
			name:        "Blank Go file in directory",
			srcPath:     `testdata/blankfile/notblank.go`,
			wantNoTests: true,
			wantErr:     true,
		}, {
			name:        "Test file with garbage data",
			srcPath:     `testdata/invalidtest/invalid.go`,
			wantNoTests: true,
			wantErr:     true,
		}, {
			name:        "Hidden file",
			srcPath:     `testdata/.hidden.go`,
			wantNoTests: true,
			wantErr:     true,
		}, {
			name:        "Nonexistant file",
			srcPath:     `testdata/nonexistant.go`,
			wantNoTests: true,
			wantErr:     true,
		}, {
			name:        "Target test file",
			srcPath:     `testdata/test100_test.go`,
			wantNoTests: true,
			wantErr:     true,
		}, {
			name:        "No funcs",
			srcPath:     `testdata/test000.go`,
			wantNoTests: true,
		}, {
			name:    "Function with neither receiver, parameters, nor results",
			srcPath: `testdata/test001.go`,
			want:    mustReadFile(t, "testdata/goldens/function_with_neither_receiver_parameters_nor_results.go"),
		}, {
			name:    "Function with anonymous arguments",
			srcPath: `testdata/test002.go`,
			want:    mustReadFile(t, "testdata/goldens/function_with_anonymous_arguments.go"),
		}, {
			name:    "Function with named argument",
			srcPath: `testdata/test003.go`,
			want:    mustReadFile(t, "testdata/goldens/function_with_named_argument.go"),
		}, {
			name:    "Function with return value",
			srcPath: `testdata/test004.go`,
			want:    mustReadFile(t, "testdata/goldens/function_with_return_value.go"),
		}, {
			name:    "Function returning an error",
			srcPath: `testdata/test005.go`,
			want:    mustReadFile(t, "testdata/goldens/function_returning_an_error.go"),
		}, {
			name:    "Function with multiple arguments",
			srcPath: `testdata/test006.go`,
			want:    mustReadFile(t, "testdata/goldens/function_with_multiple_arguments.go"),
		}, {
			name:        "Print inputs with multiple arguments",
			srcPath:     `testdata/test006.go`,
			printInputs: true,
			want:        mustReadFile(t, "testdata/goldens/print_inputs_with_multiple_arguments.go"),
		}, {
			name:    "Method on a struct pointer",
			srcPath: `testdata/test007.go`,
			want:    mustReadFile(t, "testdata/goldens/method_on_a_struct_pointer.go"),
		}, {
			name:        "Print inputs with single argument",
			srcPath:     `testdata/test007.go`,
			printInputs: true,
			want:        mustReadFile(t, "testdata/goldens/print_inputs_with_single_argument.go"),
		}, {
			name:    "Function with struct pointer argument and return type",
			srcPath: `testdata/test008.go`,
			want:    mustReadFile(t, "testdata/goldens/function_with_struct_pointer_argument_and_return_type.go"),
		}, {
			name:    "Struct value method with struct value return type",
			srcPath: `testdata/test009.go`,
			want:    mustReadFile(t, "testdata/goldens/struct_value_method_with_struct_value_return_type.go"),
		}, {
			name:    "Function with map argument and return type",
			srcPath: `testdata/test010.go`,
			want:    mustReadFile(t, "testdata/goldens/function_with_map_argument_and_return_type.go"),
		}, {
			name:    "Function with slice argument and return type",
			srcPath: `testdata/test011.go`,
			want:    mustReadFile(t, "testdata/goldens/function_with_slice_argument_and_return_type.go"),
		}, {
			name:    "Function returning only an error",
			srcPath: `testdata/test012.go`,
			want:    mustReadFile(t, "testdata/goldens/function_returning_only_an_error.go"),
		}, {
			name:    "Function with a function parameter",
			srcPath: `testdata/test013.go`,
			want:    mustReadFile(t, "testdata/goldens/function_with_a_function_parameter.go"),
		}, {
			name:    "Function with a function parameter with its own parameters and result",
			srcPath: `testdata/test014.go`,
			want:    mustReadFile(t, "testdata/goldens/function_with_a_function_parameter_with_its_own_parameters_and_result.go"),
		}, {
			name:    "Function with a function parameter that returns two results",
			srcPath: `testdata/test015.go`,
			want:    mustReadFile(t, "testdata/goldens/function_with_a_function_parameter_that_returns_two_results.go"),
		}, {
			name:    "Function with defined interface type parameter and result",
			srcPath: `testdata/test016.go`,
			want:    mustReadFile(t, "testdata/goldens/function_with_defined_interface_type_parameter_and_result.go"),
		}, {
			name:    "Function with imported interface receiver, parameter, and result",
			srcPath: `testdata/test017.go`,
			want:    mustReadFile(t, "testdata/goldens/function_with_imported_interface_receiver_parameter_and_result.go"),
		}, {
			name:    "Function with imported struct receiver, parameter, and result",
			srcPath: `testdata/test018.go`,
			want:    mustReadFile(t, "testdata/goldens/function_with_imported_struct_receiver_parameter_and_result.go"),
		}, {
			name:    "Function with multiple parameters of the same type",
			srcPath: `testdata/test019.go`,
			want:    mustReadFile(t, "testdata/goldens/function_with_multiple_parameters_of_the_same_type.go"),
		}, {
			name:    "Function with a variadic parameter",
			srcPath: `testdata/test020.go`,
			want:    mustReadFile(t, "testdata/goldens/function_with_a_variadic_parameter.go"),
		}, {
			name:    "Function with interface{} parameter and result",
			srcPath: `testdata/test021.go`,
			want:    mustReadFile(t, "testdata/goldens/function_with_interface_parameter_and_result.go"),
		}, {
			name:    "Function with named imports",
			srcPath: `testdata/test022.go`,
			want:    mustReadFile(t, "testdata/goldens/function_with_named_imports.go"),
		}, {
			name:    "Function with channel parameter and result",
			srcPath: `testdata/test023.go`,
			want:    mustReadFile(t, "testdata/goldens/function_with_channel_parameter_and_result.go"),
		}, {
			name:    "File with multiple imports",
			srcPath: `testdata/test024.go`,
			want:    mustReadFile(t, "testdata/goldens/file_with_multiple_imports.go"),
		}, {
			name:    "Function returning two results and an error",
			srcPath: `testdata/test025.go`,
			want:    mustReadFile(t, "testdata/goldens/function_returning_two_results_and_an_error.go"),
		}, {
			name:    "Multiple named results",
			srcPath: `testdata/test026.go`,
			want:    mustReadFile(t, "testdata/goldens/multiple_named_results.go"),
		}, {
			name:    "Two different structs with same method name",
			srcPath: `testdata/test027.go`,
			want:    mustReadFile(t, "testdata/goldens/two_different_structs_with_same_method_name.go"),
		}, {
			name:    "Underlying types",
			srcPath: `testdata/test028.go`,
			want:    mustReadFile(t, "testdata/goldens/underlying_types.go"),
		}, {
			name:    "Struct receiver with multiple fields",
			srcPath: `testdata/test029.go`,
			want:    mustReadFile(t, "testdata/goldens/struct_receiver_with_multiple_fields.go"),
		}, {
			name:    "Struct receiver with anonymous fields",
			srcPath: `testdata/test030.go`,
			want:    mustReadFile(t, "testdata/goldens/struct_receiver_with_anonymous_fields.go"),
		}, {
			name:    "io.Writer parameters",
			srcPath: `testdata/test031.go`,
			want:    mustReadFile(t, "testdata/goldens/io_writer_parameters.go"),
		}, {
			name:    "Two structs with same method name",
			srcPath: `testdata/test032.go`,
			want:    mustReadFile(t, "testdata/goldens/two_structs_with_same_method_name.go"),
		}, {
			name:    "Functions and methods with 'name' receivers, parameters, and results",
			srcPath: `testdata/test033.go`,
			want:    mustReadFile(t, "testdata/goldens/functions_and_methods_with_name_receivers_parameters_and_results.go"),
		}, {
			name:    "Receiver struct with reserved field names",
			srcPath: `testdata/test034.go`,
			want:    mustReadFile(t, "testdata/goldens/receiver_struct_with_reserved_field_names.go"),
		}, {
			name:    "Receiver struct with fields with complex package names",
			srcPath: `testdata/test035.go`,
			want:    mustReadFile(t, "testdata/goldens/receiver_struct_with_fields_with_complex_package_names.go"),
		}, {
			name:    "Multiple functions",
			srcPath: `testdata/test_filter.go`,
			want:    mustReadFile(t, "testdata/goldens/multiple_functions.go"),
		}, {
			name:    "Multiple functions with only",
			srcPath: `testdata/test_filter.go`,
			only:    regexp.MustCompile("FooFilter|bazFilter"),
			want:    mustReadFile(t, "testdata/goldens/multiple_functions_with_only.go"),
		}, {
			name:        "Multiple functions with only regexp without matches",
			srcPath:     `testdata/test_filter.go`,
			only:        regexp.MustCompile("asdf"),
			wantNoTests: true,
		}, {
			name:    "Multiple functions with case-insensitive only",
			srcPath: `testdata/test_filter.go`,
			only:    regexp.MustCompile("(?i)fooFilter|BazFilter"),
			want:    mustReadFile(t, "testdata/goldens/multiple_functions_with_case-insensitive_only.go"),
		}, {
			name:    "Multiple functions with only filtering on receiver",
			srcPath: `testdata/test_filter.go`,
			only:    regexp.MustCompile("^BarBarFilter$"),
			want:    mustReadFile(t, "testdata/goldens/multiple_functions_with_only_filtering_on_receiver.go"),
		}, {
			name:    "Multiple functions with only filtering on method",
			srcPath: `testdata/test_filter.go`,
			only:    regexp.MustCompile("^(BarFilter)$"),
			want:    mustReadFile(t, "testdata/goldens/multiple_functions_with_only_filtering_on_method.go"),
		}, {
			name:     "Multiple functions filtering exported",
			srcPath:  `testdata/test_filter.go`,
			exported: true,
			want:     mustReadFile(t, "testdata/goldens/multiple_functions_filtering_exported.go"),
		}, {
			name:     "Multiple functions filtering exported with only",
			srcPath:  `testdata/test_filter.go`,
			only:     regexp.MustCompile(`FooFilter`),
			exported: true,
			want:     mustReadFile(t, "testdata/goldens/multiple_functions_filtering_exported_with_only.go"),
		}, {
			name:        "Multiple functions filtering all out",
			srcPath:     `testdata/test_filter.go`,
			only:        regexp.MustCompile("fooFilter"),
			wantNoTests: true,
		}, {
			name:    "Multiple functions with excl",
			srcPath: `testdata/test_filter.go`,
			excl:    regexp.MustCompile("FooFilter|bazFilter"),
			want:    mustReadFile(t, "testdata/goldens/multiple_functions_with_excl.go"),
		}, {
			name:    "Multiple functions with case-insensitive excl",
			srcPath: `testdata/test_filter.go`,
			excl:    regexp.MustCompile("(?i)foOFilter|BaZFilter"),
			want:    mustReadFile(t, "testdata/goldens/multiple_functions_with_case-insensitive_excl.go"),
		}, {
			name:     "Multiple functions filtering exported with excl",
			srcPath:  `testdata/test_filter.go`,
			excl:     regexp.MustCompile(`FooFilter`),
			exported: true,
			want:     mustReadFile(t, "testdata/goldens/multiple_functions_filtering_exported_with_excl.go"),
		}, {
			name:        "Multiple functions excluding all",
			srcPath:     `testdata/test_filter.go`,
			excl:        regexp.MustCompile("bazFilter|FooFilter|BarFilter"),
			wantNoTests: true,
		}, {
			name:    "Multiple functions excluding on receiver",
			srcPath: `testdata/test_filter.go`,
			excl:    regexp.MustCompile("^BarBarFilter$"),
			want:    mustReadFile(t, "testdata/goldens/multiple_functions_excluding_on_receiver.go"),
		}, {
			name:    "Multiple functions excluding on method",
			srcPath: `testdata/test_filter.go`,
			excl:    regexp.MustCompile("^BarFilter$"),
			want:    mustReadFile(t, "testdata/goldens/multiple_functions_excluding_on_method.go"),
		}, {
			name:    "Multiple functions with both only and excl",
			srcPath: `testdata/test_filter.go`,
			only:    regexp.MustCompile("BarFilter"),
			excl:    regexp.MustCompile("FooFilter"),
			want:    mustReadFile(t, "testdata/goldens/multiple_functions_with_both_only_and_excl.go"),
		}, {
			name:    "Multiple functions with only and excl competing",
			srcPath: `testdata/test_filter.go`,
			only:    regexp.MustCompile("FooFilter|BarFilter"),
			excl:    regexp.MustCompile("FooFilter|bazFilter"),
			want:    mustReadFile(t, "testdata/goldens/multiple_functions_with_only_and_excl_competing.go"),
		}, {
			name:     "Custom importer fails",
			srcPath:  `testdata/test_filter.go`,
			importer: &fakeImporter{err: errors.New("error")},
			want:     mustReadFile(t, "testdata/goldens/custom_importer_fails.go"),
		}, {
			name:    "Existing test file",
			srcPath: `testdata/test100.go`,
			want:    mustReadFile(t, "testdata/goldens/existing_test_file.go"),
		}, {
			name:    "Existing test file with just package declaration",
			srcPath: `testdata/test101.go`,
			want:    mustReadFile(t, "testdata/goldens/existing_test_file_with_just_package_declaration.go"),
		}, {
			name:    "Existing test file with no functions",
			srcPath: `testdata/test102.go`,
			want:    mustReadFile(t, "testdata/goldens/existing_test_file_with_no_functions.go"),
		}, {
			name:    "Existing test file with multiple imports",
			srcPath: `testdata/test200.go`,
			want:    mustReadFile(t, "testdata/goldens/existing_test_file_with_multiple_imports.go"),
		}, {
			name:              "Entire testdata directory",
			srcPath:           `testdata/`,
			wantMultipleTests: true,
		}, {
			name:    "Different packages in same directory - part 1",
			srcPath: `testdata/mixedpkg/bar.go`,
			want:    mustReadFile(t, "testdata/goldens/different_packages_in_same_directory_-_part_1.go"),
		}, {
			name:    "Different packages in same directory - part 2",
			srcPath: `testdata/mixedpkg/foo.go`,
			want:    mustReadFile(t, "testdata/goldens/different_packages_in_same_directory_-_part_2.go"),
		}, {
			name:    "Empty test file",
			srcPath: `testdata/blanktest/blank.go`,
			want:    mustReadFile(t, "testdata/goldens/empty_test_file.go"),
		}, {
			name:    "Test file with syntax errors",
			srcPath: `testdata/syntaxtest/syntax.go`,
			want:    mustReadFile(t, "testdata/goldens/test_file_with_syntax_errors.go"),
		}, {
			name:    "Undefined types",
			srcPath: `testdata/undefinedtypes/undefined.go`,
			want:    mustReadFile(t, "testdata/goldens/undefined_types.go"),
		},
	}
	tmp, err := ioutil.TempDir("", "gotests_test")
	if err != nil {
		t.Fatalf("ioutil.TempDir: %v", err)
	}
	for _, tt := range tests {
		gts, err := GenerateTests(tt.srcPath, &Options{
			Only:        tt.only,
			Exclude:     tt.excl,
			Exported:    tt.exported,
			PrintInputs: tt.printInputs,
			Importer:    func() types.Importer { return tt.importer },
		})
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. GenerateTests(%v) error = %v, wantErr %v", tt.name, tt.srcPath, err, tt.wantErr)
			continue
		}
		if (len(gts) == 0) != tt.wantNoTests {
			t.Errorf("%q. GenerateTests(%v) returned no tests", tt.name, tt.srcPath)
			continue
		}
		if (len(gts) > 1) != tt.wantMultipleTests {
			t.Errorf("%q. GenerateTests(%v) returned too many tests", tt.name, tt.srcPath)
			continue
		}
		if tt.wantNoTests || tt.wantMultipleTests {
			continue
		}
		if got := string(gts[0].Output); got != tt.want {
			t.Errorf("%q. GenerateTests(%v) = \n%v, want \n%v", tt.name, tt.srcPath, got, tt.want)
			outputResult(t, tmp, tt.name, gts[0].Output)
		}
	}
}

func mustReadFile(t *testing.T, filename string) string {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		t.Fatal(err)
	}
	return string(b)
}

func outputResult(t *testing.T, tmpDir, testName string, got []byte) {
	tmpResult := path.Join(tmpDir, toSnakeCase(testName)+".go")
	if err := ioutil.WriteFile(tmpResult, got, 0644); err != nil {
		t.Errorf("ioutil.WriteFile: %v", err)
	}
	t.Logf(tmpResult)
}

func toSnakeCase(s string) string {
	var res []rune
	for _, r := range []rune(s) {
		r = unicode.ToLower(r)
		switch r {
		case ' ', '.':
			r = '_'
		case ',', '\'', '{', '}':
			continue
		}
		res = append(res, r)
	}
	return string(res)
}

// 249032394 ns/op
func BenchmarkGenerateTests(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenerateTests("testdata/", &Options{})
	}
}

// A fake importer.
type fakeImporter struct {
	err error
}

func (f *fakeImporter) Import(path string) (*types.Package, error) {
	return nil, f.err
}
